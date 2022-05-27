package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Gnmi struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Delete     types.Bool   `tfsdk:"delete"`
	Attributes types.Map    `tfsdk:"attributes"`
	Lists      []GnmiList   `tfsdk:"lists"`
}

type GnmiList struct {
	Name  types.String `tfsdk:"name"`
	Key   types.String `tfsdk:"key"`
	Items []types.Map  `tfsdk:"items"`
}

type GnmiDataSource struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Attributes types.Map    `tfsdk:"attributes"`
}

func (data Gnmi) getPath() string {
	return data.Path.Value
}

func (data Gnmi) toBody(ctx context.Context) string {
	body := "{}"

	var attributes map[string]string
	data.Attributes.ElementsAs(ctx, &attributes, false)

	for attr, value := range attributes {
		attr = strings.ReplaceAll(attr, "/", ".")
		body, _ = sjson.Set(body, attr, value)
	}

	for i := range data.Lists {
		listName := strings.ReplaceAll(data.Lists[i].Name.Value, "/", ".")
		body, _ = sjson.Set(body, listName, []interface{}{})
		for ii := range data.Lists[i].Items {
			var listAttributes map[string]string
			data.Lists[i].Items[ii].ElementsAs(ctx, &listAttributes, false)
			attrs := ""
			for attr, value := range listAttributes {
				attr = strings.ReplaceAll(attr, "/", ".")
				attrs, _ = sjson.Set(attrs, attr, value)
			}
			body, _ = sjson.SetRaw(body, listName+".-1", attrs)
		}
	}

	return body
}

func (data *Gnmi) fromBody(ctx context.Context, res []byte) {
	for attr := range data.Attributes.Elems {
		attrPath := strings.ReplaceAll(attr, "/", ".")
		value := gjson.GetBytes(res, attrPath)
		if !value.Exists() ||
			(value.IsObject() && len(value.Map()) == 0) ||
			value.Raw == "[null]" {

			data.Attributes.Elems[attr] = types.String{Value: ""}
		} else {
			data.Attributes.Elems[attr] = types.String{Value: value.String()}
		}
	}

	for i := range data.Lists {
		keys := strings.Split(data.Lists[i].Key.Value, ",")
		for ii := range data.Lists[i].Items {
			var keyValues []string
			for _, key := range keys {
				v, _ := data.Lists[i].Items[ii].Elems[key].ToTerraformValue(ctx)
				var keyValue string
				v.As(&keyValue)
				keyValues = append(keyValues, keyValue)
			}

			// find item by key(s)
			var r gjson.Result
			namePath := strings.ReplaceAll(data.Lists[i].Name.Value, "/", ".")
			gjson.GetBytes(res, namePath).ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() == keyValues[ik] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						r = v
						return false
					}
					return true
				},
			)

			for attr := range data.Lists[i].Items[ii].Elems {
				attrPath := strings.ReplaceAll(attr, "/", ".")
				value := r.Get(attrPath)
				if !value.Exists() ||
					(value.IsObject() && len(value.Map()) == 0) ||
					value.Raw == "[null]" {

					data.Lists[i].Items[ii].Elems[attr] = types.String{Value: ""}
				} else {
					data.Lists[i].Items[ii].Elems[attr] = types.String{Value: value.String()}
				}
			}
		}
	}
}

func (data *Gnmi) getDeletedListItems(ctx context.Context, state Gnmi) []string {
	deletedListItems := make([]string, 0)
	for l := range state.Lists {
		name := state.Lists[l].Name.Value
		keys := strings.Split(state.Lists[l].Key.Value, ",")
		var dataList GnmiList
		for _, dl := range data.Lists {
			if dl.Name.Value == name {
				dataList = dl
			}
		}
		// check if state item is also included in plan, if not delete item
		for i := range state.Lists[l].Items {
			var slia map[string]string
			state.Lists[l].Items[i].ElementsAs(ctx, &slia, false)

			// if state key values are empty move on to next item
			emptyKey := false
			for _, key := range keys {
				if slia[key] == "" {
					emptyKey = true
					break
				}
			}
			if emptyKey {
				continue
			}

			// find data (plan) item with matching key values
			found := false
			for dli := range dataList.Items {
				var dlia map[string]string
				dataList.Items[dli].ElementsAs(ctx, &dlia, false)
				for _, key := range keys {
					if dlia[key] == slia[key] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					break
				}
			}

			// if no matching item in plan found -> delete
			if !found {
				keyString := ""
				for _, key := range keys {
					keyString += fmt.Sprintf("[%s=%s]", key, slia[key])
				}
				deletedListItems = append(deletedListItems, state.getPath()+"/"+name+keyString)
			}
		}
	}
	return deletedListItems
}
