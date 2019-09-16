for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        env GOOS=$GOOS GOARCH=$GOARCH go build -v -o ./build/vc-deploy-$GOOS-$GOARCH
    done
done
