cd proto
buf generate
buf generate --template buf.gen.pulsar.yaml
cd ..

cp -r github.com/circlefin/noble-fiattokenfactory/* ./
rm -rf circle
rm -rf github.com
