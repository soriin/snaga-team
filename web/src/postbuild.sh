if [ ! -d "../public/bower_components" ]; then
  mkdir -p "../public/bower_components"
fi
cp -R  bower_components/*paper* ../public/bower_components/
cp -R  bower_components/*iron* ../public/bower_components/
cp -R  bower_components/*polymer* ../public/bower_components/
