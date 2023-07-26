generate_ent:
    go generate ./ent

_yaml_merge output +FILES:
    #!/usr/bin/env bash
    yq ea '. as $item ireduce ({}; . * $item )' {{FILES}} | tee {{output}} >/dev/null

api_docs:
    #!/usr/bin/env bash
    GLOBIGNORE='docs/manekani/apis/*.openapi.*'
    for f in docs/manekani/apis/*.yaml; do
        just _yaml_merge `dirname $f`/`basename -s .yaml $f`.openapi.yaml docs/manekani/base.yaml "$f"
    done
    just _yaml_merge docs/manekani/openapi.yaml docs/manekani/base.yaml docs/manekani/apis/*.yaml

api_interfaces: api_docs
    #!/usr/bin/env bash
    shopt -s extglob
    for f in docs/manekani/apis/*.!(yaml); do
        PKGNAME=`basename -s .openapi.yaml $f`
        oapi-codegen -package "$PKGNAME" \
            -generate gin,types,spec \
            "$f" > ./server/api/"$PKGNAME"/api.gen.go
    done


start:
    docker compose up -d