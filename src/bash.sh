./gomobile build -o=app.apk -target=android .

rm  -r lib 
unzip  app.apk   'lib*'

cp lib/arm64-v8a/libapp.so ../example/android/libs/arm64-v8a/libexample.so
cp lib/armeabi-v7a/libapp.so ../example/android/libs/armeabi-v7a/libexample.so
cp lib/x86/libapp.so ../example/android/libs/x86/libexample.so
cp lib/x86_64/libapp.so ../example/android/libs/x86_64/libexample.so

cp -r assets/ ../example/android/assets/

cd ../example
./gradlew 
./gradlew assembleDebug

cp  ./android/build/outputs/apk/debug/android-debug.apk ../src/out


clear
echo "Ok"
