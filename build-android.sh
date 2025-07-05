#!/bin/bash
set -e

NDK_PATH=$HOME/library/Android/Sdk/ndk/28.0.12916984
API=21
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
OUTPUT_DIR="$SCRIPT_DIR/libs/android"
JNIWRAPPER="$SCRIPT_DIR/jniwrapper.go"

mkdir -p "$OUTPUT_DIR"

ARCHS=("armeabi-v7a" "arm64-v8a" "x86" "x86_64")

for ABI in "${ARCHS[@]}"; do
  echo "🔧 Building for $ABI..."

  case $ABI in
    "armeabi-v7a")
      export GOARCH=arm
      export GOARM=7
      TARGET_HOST=armv7a-linux-androideabi
      ;;
    "arm64-v8a")
      export GOARCH=arm64
      TARGET_HOST=aarch64-linux-android
      ;;
    "x86")
      export GOARCH=386
      TARGET_HOST=i686-linux-android
      ;;
    "x86_64")
      export GOARCH=amd64
      TARGET_HOST=x86_64-linux-android
      ;;
  esac

  export GOOS=android
  export CC=$NDK_PATH/toolchains/llvm/prebuilt/darwin-x86_64/bin/${TARGET_HOST}${API}-clang
  export CGO_ENABLED=1
  OUT_DIR="$OUTPUT_DIR/$ABI"
  mkdir -p "$OUT_DIR"

  echo "🏗️ go build -buildmode=c-shared -o $OUT_DIR/libsecretr.so $JNIWRAPPER"
  go build  -buildmode=c-shared -o "$OUT_DIR/libsecretr.so" "$JNIWRAPPER"
done

echo "✅ All builds complete! Output in: $OUTPUT_DIR"
