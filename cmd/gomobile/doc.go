// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by 'gomobile help documentation doc.go'. DO NOT EDIT.

/*
Gomobile is a tool for building and running mobile apps written in Go.

To install:

	$ go get github.com/sumup/mobile/cmd/gomobile
	$ gomobile init

At least Go 1.10 is required.
For detailed instructions, see https://golang.org/wiki/Mobile.

Usage:

	gomobile command [arguments]

Commands:

	bind        build a library for Android and iOS
	build       compile android APK and iOS app
	clean       remove object files and cached gomobile files
	init        install NDK toolchains and build OpenAL for Android
	install     compile android APK and install on device
	version     print version

Use 'gomobile help [command]' for more information about that command.


Build a library for Android and iOS

Usage:

	gomobile bind [-target android|ios] [-bootclasspath <path>] [-classpath <path>] [-o output] [build flags] [package]

Bind generates language bindings for the package named by the import
path, and compiles a library for the named target system.

The -target flag takes a target system name, either android (the
default) or ios.

For -target android, the bind command produces an AAR (Android ARchive)
file that archives the precompiled Java API stub classes, the compiled
shared libraries, and all asset files in the /assets subdirectory under
the package directory. The output is named '<package_name>.aar' by
default. This AAR file is commonly used for binary distribution of an
Android library project and most Android IDEs support AAR import. For
example, in Android Studio (1.2+), an AAR file can be imported using
the module import wizard (File > New > New Module > Import .JAR or
.AAR package), and setting it as a new dependency
(File > Project Structure > Dependencies).  This requires 'javac'
(version 1.7+) and Android SDK (API level 15 or newer) to build the
library for Android. The environment variable ANDROID_HOME must be set
to the path to Android SDK. Use the -javapkg flag to specify the Java
package prefix for the generated classes.

By default, -target=android builds shared libraries for all supported
instruction sets (arm, arm64, 386, amd64). A subset of instruction sets
can be selected by specifying target type with the architecture name. E.g.,
-target=android/arm,android/386.

For -target ios, gomobile must be run on an OS X machine with Xcode
installed. The generated Objective-C types can be prefixed with the -prefix
flag.

For -target android, the -bootclasspath and -classpath flags are used to
control the bootstrap classpath and the classpath for Go wrappers to Java
classes.

The -v flag provides verbose output, including the list of packages built.

The build flags -a, -n, -x, -gcflags, -ldflags, -tags, and -work
are shared with the build command. For documentation, see 'go help build'.


Compile android APK and iOS app

Usage:

	gomobile build [-target android|ios] [-o output] [-bundleid bundleID] [build flags] [package]

Build compiles and encodes the app named by the import path.

The named package must define a main function.

The -target flag takes a target system name, either android (the
default) or ios.

For -target android, if an AndroidManifest.xml is defined in the
package directory, it is added to the APK output. Otherwise, a default
manifest is generated. By default, this builds a fat APK for all supported
instruction sets (arm, 386, amd64, arm64). A subset of instruction sets can
be selected by specifying target type with the architecture name. E.g.
-target=android/arm,android/386.

For -target ios, gomobile must be run on an OS X machine with Xcode
installed.

If the package directory contains an assets subdirectory, its contents
are copied into the output.

Flag -iosversion sets the minimal version of the iOS SDK to compile against.
The default version is 6.1.

The -bundleid flag is required for -target ios and sets the bundle ID to use
with the app.

The -o flag specifies the output file name. If not specified, the
output file name depends on the package built.

The -v flag provides verbose output, including the list of packages built.

The build flags -a, -i, -n, -x, -gcflags, -ldflags, -tags, and -work are
shared with the build command. For documentation, see 'go help build'.


Remove object files and cached gomobile files

Usage:

	gomobile clean

Clean removes object files and cached NDK files downloaded by gomobile init


Install NDK toolchains and build OpenAL for Android

Usage:

	gomobile init [-ndk dir] [-openal dir]

If the -ndk flag is specified or the Android NDK is installed at
$ANDROID_HOME/ndk-bundle, init will create NDK standalone toolchains
for Android targets.

If a OpenAL source directory is specified with -openal, init will
build an Android version of OpenAL for use with gomobile build
and gomobile install.


Compile android APK and install on device

Usage:

	gomobile install [-target android] [build flags] [package]

Install compiles and installs the app named by the import path on the
attached mobile device.

Only -target android is supported. The 'adb' tool must be on the PATH.

The build flags -a, -i, -n, -x, -gcflags, -ldflags, -tags, and -work are
shared with the build command.
For documentation, see 'go help build'.


Print version

Usage:

	gomobile version

Version prints versions of the gomobile binary and tools
*/
package main // import "github.com/sumup/mobile/cmd/gomobile"
