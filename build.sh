
gox -output "dist/{{.OS}}_{{.Arch}}/{{.Dir}}-{{.OS}}-{{.Arch}}" \
    ./apps/...


for f in ./dist/*; do
    if [ -d ${f} ]; then
        # Will not run if no directories are available
        echo $f
		cp ./data/* $f
    fi
done