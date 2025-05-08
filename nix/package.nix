{ buildGoModule }:

buildGoModule {
  pname = "cgroup-exporter";
  version = "0.1.0";
  src = ../.;
  vendorHash = "sha256-f0f8tYmoI6DtuB/K4++gu9b2na/d0ECTaF2zvDijW58=";
}
