resource "iosxr_sampler_map" "example" {
  sampler_map_name = "sampler_map1"
  random           = 1
  out_of           = 1
}
