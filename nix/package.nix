{ buildGoModule }:

buildGoModule {
  pname = "cgroup-exporter";
  version = "0.1.0";
  src = ../.;
  vendorHash = "sha256-B9ppgJQout7eabd58iAAlgELRM4UFcNVhy50Hokhras=";
}
