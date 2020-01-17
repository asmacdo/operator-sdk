# TODO(asmacdo) dont install and build in here, just run liche
pushd website
npm install postcss-cli
hugo
liche -d public -r -c 50 public
popd
