resource "iosxr_flow_sampler_map" "example" {
  name   = "sampler_map1"
  random = 1
  out_of = 1
}
