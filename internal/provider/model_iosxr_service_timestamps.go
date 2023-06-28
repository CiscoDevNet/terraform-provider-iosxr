// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ServiceTimestamps struct {
	Device                     types.String `tfsdk:"device"`
	Id                         types.String `tfsdk:"id"`
	DebugDatetimeLocaltimeOnly types.Bool   `tfsdk:"debug_datetime_localtime_only"`
	DebugDatetimeLocaltime     types.Bool   `tfsdk:"debug_datetime_localtime"`
	DebugDatetimeMsec          types.Bool   `tfsdk:"debug_datetime_msec"`
	DebugDatetimeShowTimezone  types.Bool   `tfsdk:"debug_datetime_show_timezone"`
	DebugDatetimeYear          types.Bool   `tfsdk:"debug_datetime_year"`
	DebugUptime                types.Bool   `tfsdk:"debug_uptime"`
	DebugDisable               types.Bool   `tfsdk:"debug_disable"`
	LogDatetimeLocaltimeOnly   types.Bool   `tfsdk:"log_datetime_localtime_only"`
	LogDatetimeLocaltime       types.Bool   `tfsdk:"log_datetime_localtime"`
	LogDatetimeMsec            types.Bool   `tfsdk:"log_datetime_msec"`
	LogDatetimeShowTimezone    types.Bool   `tfsdk:"log_datetime_show_timezone"`
	LogDatetimeYear            types.Bool   `tfsdk:"log_datetime_year"`
	LogUptime                  types.Bool   `tfsdk:"log_uptime"`
	LogDisable                 types.Bool   `tfsdk:"log_disable"`
}

type ServiceTimestampsData struct {
	Device                     types.String `tfsdk:"device"`
	Id                         types.String `tfsdk:"id"`
	DebugDatetimeLocaltimeOnly types.Bool   `tfsdk:"debug_datetime_localtime_only"`
	DebugDatetimeLocaltime     types.Bool   `tfsdk:"debug_datetime_localtime"`
	DebugDatetimeMsec          types.Bool   `tfsdk:"debug_datetime_msec"`
	DebugDatetimeShowTimezone  types.Bool   `tfsdk:"debug_datetime_show_timezone"`
	DebugDatetimeYear          types.Bool   `tfsdk:"debug_datetime_year"`
	DebugUptime                types.Bool   `tfsdk:"debug_uptime"`
	DebugDisable               types.Bool   `tfsdk:"debug_disable"`
	LogDatetimeLocaltimeOnly   types.Bool   `tfsdk:"log_datetime_localtime_only"`
	LogDatetimeLocaltime       types.Bool   `tfsdk:"log_datetime_localtime"`
	LogDatetimeMsec            types.Bool   `tfsdk:"log_datetime_msec"`
	LogDatetimeShowTimezone    types.Bool   `tfsdk:"log_datetime_show_timezone"`
	LogDatetimeYear            types.Bool   `tfsdk:"log_datetime_year"`
	LogUptime                  types.Bool   `tfsdk:"log_uptime"`
	LogDisable                 types.Bool   `tfsdk:"log_disable"`
}

func (data ServiceTimestamps) getPath() string {
	return "Cisco-IOS-XR-um-service-timestamps-cfg:/service/timestamps"
}

func (data ServiceTimestampsData) getPath() string {
	return "Cisco-IOS-XR-um-service-timestamps-cfg:/service/timestamps"
}

func (data ServiceTimestamps) toBody(ctx context.Context) string {
	body := "{}"
	if !data.DebugDatetimeLocaltimeOnly.IsNull() && !data.DebugDatetimeLocaltimeOnly.IsUnknown() {
		if data.DebugDatetimeLocaltimeOnly.ValueBool() {
			body, _ = sjson.Set(body, "debug.datetime.localtime-only", map[string]string{})
		}
	}
	if !data.DebugDatetimeLocaltime.IsNull() && !data.DebugDatetimeLocaltime.IsUnknown() {
		if data.DebugDatetimeLocaltime.ValueBool() {
			body, _ = sjson.Set(body, "debug.datetime.localtime", map[string]string{})
		}
	}
	if !data.DebugDatetimeMsec.IsNull() && !data.DebugDatetimeMsec.IsUnknown() {
		if data.DebugDatetimeMsec.ValueBool() {
			body, _ = sjson.Set(body, "debug.datetime.msec", map[string]string{})
		}
	}
	if !data.DebugDatetimeShowTimezone.IsNull() && !data.DebugDatetimeShowTimezone.IsUnknown() {
		if data.DebugDatetimeShowTimezone.ValueBool() {
			body, _ = sjson.Set(body, "debug.datetime.show-timezone", map[string]string{})
		}
	}
	if !data.DebugDatetimeYear.IsNull() && !data.DebugDatetimeYear.IsUnknown() {
		if data.DebugDatetimeYear.ValueBool() {
			body, _ = sjson.Set(body, "debug.datetime.year", map[string]string{})
		}
	}
	if !data.DebugUptime.IsNull() && !data.DebugUptime.IsUnknown() {
		if data.DebugUptime.ValueBool() {
			body, _ = sjson.Set(body, "debug.uptime", map[string]string{})
		}
	}
	if !data.DebugDisable.IsNull() && !data.DebugDisable.IsUnknown() {
		if data.DebugDisable.ValueBool() {
			body, _ = sjson.Set(body, "debug.disable", map[string]string{})
		}
	}
	if !data.LogDatetimeLocaltimeOnly.IsNull() && !data.LogDatetimeLocaltimeOnly.IsUnknown() {
		if data.LogDatetimeLocaltimeOnly.ValueBool() {
			body, _ = sjson.Set(body, "log.datetime.localtime-only", map[string]string{})
		}
	}
	if !data.LogDatetimeLocaltime.IsNull() && !data.LogDatetimeLocaltime.IsUnknown() {
		if data.LogDatetimeLocaltime.ValueBool() {
			body, _ = sjson.Set(body, "log.datetime.localtime", map[string]string{})
		}
	}
	if !data.LogDatetimeMsec.IsNull() && !data.LogDatetimeMsec.IsUnknown() {
		if data.LogDatetimeMsec.ValueBool() {
			body, _ = sjson.Set(body, "log.datetime.msec", map[string]string{})
		}
	}
	if !data.LogDatetimeShowTimezone.IsNull() && !data.LogDatetimeShowTimezone.IsUnknown() {
		if data.LogDatetimeShowTimezone.ValueBool() {
			body, _ = sjson.Set(body, "log.datetime.show-timezone", map[string]string{})
		}
	}
	if !data.LogDatetimeYear.IsNull() && !data.LogDatetimeYear.IsUnknown() {
		if data.LogDatetimeYear.ValueBool() {
			body, _ = sjson.Set(body, "log.datetime.year", map[string]string{})
		}
	}
	if !data.LogUptime.IsNull() && !data.LogUptime.IsUnknown() {
		if data.LogUptime.ValueBool() {
			body, _ = sjson.Set(body, "log.uptime", map[string]string{})
		}
	}
	if !data.LogDisable.IsNull() && !data.LogDisable.IsUnknown() {
		if data.LogDisable.ValueBool() {
			body, _ = sjson.Set(body, "log.disable", map[string]string{})
		}
	}
	return body
}

func (data *ServiceTimestamps) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "debug.datetime.localtime-only"); !data.DebugDatetimeLocaltimeOnly.IsNull() {
		if value.Exists() {
			data.DebugDatetimeLocaltimeOnly = types.BoolValue(true)
		} else {
			data.DebugDatetimeLocaltimeOnly = types.BoolValue(false)
		}
	} else {
		data.DebugDatetimeLocaltimeOnly = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "debug.datetime.localtime"); !data.DebugDatetimeLocaltime.IsNull() {
		if value.Exists() {
			data.DebugDatetimeLocaltime = types.BoolValue(true)
		} else {
			data.DebugDatetimeLocaltime = types.BoolValue(false)
		}
	} else {
		data.DebugDatetimeLocaltime = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "debug.datetime.msec"); !data.DebugDatetimeMsec.IsNull() {
		if value.Exists() {
			data.DebugDatetimeMsec = types.BoolValue(true)
		} else {
			data.DebugDatetimeMsec = types.BoolValue(false)
		}
	} else {
		data.DebugDatetimeMsec = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "debug.datetime.show-timezone"); !data.DebugDatetimeShowTimezone.IsNull() {
		if value.Exists() {
			data.DebugDatetimeShowTimezone = types.BoolValue(true)
		} else {
			data.DebugDatetimeShowTimezone = types.BoolValue(false)
		}
	} else {
		data.DebugDatetimeShowTimezone = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "debug.datetime.year"); !data.DebugDatetimeYear.IsNull() {
		if value.Exists() {
			data.DebugDatetimeYear = types.BoolValue(true)
		} else {
			data.DebugDatetimeYear = types.BoolValue(false)
		}
	} else {
		data.DebugDatetimeYear = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "debug.uptime"); !data.DebugUptime.IsNull() {
		if value.Exists() {
			data.DebugUptime = types.BoolValue(true)
		} else {
			data.DebugUptime = types.BoolValue(false)
		}
	} else {
		data.DebugUptime = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "debug.disable"); !data.DebugDisable.IsNull() {
		if value.Exists() {
			data.DebugDisable = types.BoolValue(true)
		} else {
			data.DebugDisable = types.BoolValue(false)
		}
	} else {
		data.DebugDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "log.datetime.localtime-only"); !data.LogDatetimeLocaltimeOnly.IsNull() {
		if value.Exists() {
			data.LogDatetimeLocaltimeOnly = types.BoolValue(true)
		} else {
			data.LogDatetimeLocaltimeOnly = types.BoolValue(false)
		}
	} else {
		data.LogDatetimeLocaltimeOnly = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "log.datetime.localtime"); !data.LogDatetimeLocaltime.IsNull() {
		if value.Exists() {
			data.LogDatetimeLocaltime = types.BoolValue(true)
		} else {
			data.LogDatetimeLocaltime = types.BoolValue(false)
		}
	} else {
		data.LogDatetimeLocaltime = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "log.datetime.msec"); !data.LogDatetimeMsec.IsNull() {
		if value.Exists() {
			data.LogDatetimeMsec = types.BoolValue(true)
		} else {
			data.LogDatetimeMsec = types.BoolValue(false)
		}
	} else {
		data.LogDatetimeMsec = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "log.datetime.show-timezone"); !data.LogDatetimeShowTimezone.IsNull() {
		if value.Exists() {
			data.LogDatetimeShowTimezone = types.BoolValue(true)
		} else {
			data.LogDatetimeShowTimezone = types.BoolValue(false)
		}
	} else {
		data.LogDatetimeShowTimezone = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "log.datetime.year"); !data.LogDatetimeYear.IsNull() {
		if value.Exists() {
			data.LogDatetimeYear = types.BoolValue(true)
		} else {
			data.LogDatetimeYear = types.BoolValue(false)
		}
	} else {
		data.LogDatetimeYear = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "log.uptime"); !data.LogUptime.IsNull() {
		if value.Exists() {
			data.LogUptime = types.BoolValue(true)
		} else {
			data.LogUptime = types.BoolValue(false)
		}
	} else {
		data.LogUptime = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "log.disable"); !data.LogDisable.IsNull() {
		if value.Exists() {
			data.LogDisable = types.BoolValue(true)
		} else {
			data.LogDisable = types.BoolValue(false)
		}
	} else {
		data.LogDisable = types.BoolNull()
	}
}

func (data *ServiceTimestampsData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "debug.datetime.localtime-only"); value.Exists() {
		data.DebugDatetimeLocaltimeOnly = types.BoolValue(true)
	} else {
		data.DebugDatetimeLocaltimeOnly = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "debug.datetime.localtime"); value.Exists() {
		data.DebugDatetimeLocaltime = types.BoolValue(true)
	} else {
		data.DebugDatetimeLocaltime = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "debug.datetime.msec"); value.Exists() {
		data.DebugDatetimeMsec = types.BoolValue(true)
	} else {
		data.DebugDatetimeMsec = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "debug.datetime.show-timezone"); value.Exists() {
		data.DebugDatetimeShowTimezone = types.BoolValue(true)
	} else {
		data.DebugDatetimeShowTimezone = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "debug.datetime.year"); value.Exists() {
		data.DebugDatetimeYear = types.BoolValue(true)
	} else {
		data.DebugDatetimeYear = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "debug.uptime"); value.Exists() {
		data.DebugUptime = types.BoolValue(true)
	} else {
		data.DebugUptime = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "debug.disable"); value.Exists() {
		data.DebugDisable = types.BoolValue(true)
	} else {
		data.DebugDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "log.datetime.localtime-only"); value.Exists() {
		data.LogDatetimeLocaltimeOnly = types.BoolValue(true)
	} else {
		data.LogDatetimeLocaltimeOnly = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "log.datetime.localtime"); value.Exists() {
		data.LogDatetimeLocaltime = types.BoolValue(true)
	} else {
		data.LogDatetimeLocaltime = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "log.datetime.msec"); value.Exists() {
		data.LogDatetimeMsec = types.BoolValue(true)
	} else {
		data.LogDatetimeMsec = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "log.datetime.show-timezone"); value.Exists() {
		data.LogDatetimeShowTimezone = types.BoolValue(true)
	} else {
		data.LogDatetimeShowTimezone = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "log.datetime.year"); value.Exists() {
		data.LogDatetimeYear = types.BoolValue(true)
	} else {
		data.LogDatetimeYear = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "log.uptime"); value.Exists() {
		data.LogUptime = types.BoolValue(true)
	} else {
		data.LogUptime = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "log.disable"); value.Exists() {
		data.LogDisable = types.BoolValue(true)
	} else {
		data.LogDisable = types.BoolValue(false)
	}
}

func (data *ServiceTimestamps) getDeletedListItems(ctx context.Context, state ServiceTimestamps) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *ServiceTimestamps) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.DebugDatetimeLocaltimeOnly.IsNull() && !data.DebugDatetimeLocaltimeOnly.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/debug/datetime/localtime-only", data.getPath()))
	}
	if !data.DebugDatetimeLocaltime.IsNull() && !data.DebugDatetimeLocaltime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/debug/datetime/localtime", data.getPath()))
	}
	if !data.DebugDatetimeMsec.IsNull() && !data.DebugDatetimeMsec.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/debug/datetime/msec", data.getPath()))
	}
	if !data.DebugDatetimeShowTimezone.IsNull() && !data.DebugDatetimeShowTimezone.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/debug/datetime/show-timezone", data.getPath()))
	}
	if !data.DebugDatetimeYear.IsNull() && !data.DebugDatetimeYear.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/debug/datetime/year", data.getPath()))
	}
	if !data.DebugUptime.IsNull() && !data.DebugUptime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/debug/uptime", data.getPath()))
	}
	if !data.DebugDisable.IsNull() && !data.DebugDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/debug/disable", data.getPath()))
	}
	if !data.LogDatetimeLocaltimeOnly.IsNull() && !data.LogDatetimeLocaltimeOnly.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/log/datetime/localtime-only", data.getPath()))
	}
	if !data.LogDatetimeLocaltime.IsNull() && !data.LogDatetimeLocaltime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/log/datetime/localtime", data.getPath()))
	}
	if !data.LogDatetimeMsec.IsNull() && !data.LogDatetimeMsec.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/log/datetime/msec", data.getPath()))
	}
	if !data.LogDatetimeShowTimezone.IsNull() && !data.LogDatetimeShowTimezone.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/log/datetime/show-timezone", data.getPath()))
	}
	if !data.LogDatetimeYear.IsNull() && !data.LogDatetimeYear.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/log/datetime/year", data.getPath()))
	}
	if !data.LogUptime.IsNull() && !data.LogUptime.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/log/uptime", data.getPath()))
	}
	if !data.LogDisable.IsNull() && !data.LogDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/log/disable", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *ServiceTimestamps) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.DebugDatetimeLocaltimeOnly.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/debug/datetime/localtime-only", data.getPath()))
	}
	if !data.DebugDatetimeLocaltime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/debug/datetime/localtime", data.getPath()))
	}
	if !data.DebugDatetimeMsec.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/debug/datetime/msec", data.getPath()))
	}
	if !data.DebugDatetimeShowTimezone.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/debug/datetime/show-timezone", data.getPath()))
	}
	if !data.DebugDatetimeYear.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/debug/datetime/year", data.getPath()))
	}
	if !data.DebugUptime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/debug/uptime", data.getPath()))
	}
	if !data.DebugDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/debug/disable", data.getPath()))
	}
	if !data.LogDatetimeLocaltimeOnly.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/log/datetime/localtime-only", data.getPath()))
	}
	if !data.LogDatetimeLocaltime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/log/datetime/localtime", data.getPath()))
	}
	if !data.LogDatetimeMsec.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/log/datetime/msec", data.getPath()))
	}
	if !data.LogDatetimeShowTimezone.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/log/datetime/show-timezone", data.getPath()))
	}
	if !data.LogDatetimeYear.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/log/datetime/year", data.getPath()))
	}
	if !data.LogUptime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/log/uptime", data.getPath()))
	}
	if !data.LogDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/log/disable", data.getPath()))
	}
	return deletePaths
}
