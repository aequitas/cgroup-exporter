{ buildGoModule }:

buildGoModule {
  pname = "cgroup-exporter";
  version = "0.1.0";
  src = ../.;
  vendorHash = "sha256-NnvB20rORPS5QF5enbb5KpWaKZ70ybSgfd7wjk21/Cg=";
}
