#!/usr/bin/perl

use strict;
use warnings;
use 5.018;

my $os_arch = [
    "darwin    386",
    "darwin    amd64",
    "darwin    arm",
    "darwin    arm64",
    "dragonfly amd64",
    "freebsd   386",
    "freebsd   amd64",
    "freebsd   arm",
    "linux     386",
    "linux     amd64",
    "linux     arm",
    "linux     arm64",
    "linux     ppc64",
    "linux     ppc64le",
    "linux     mips64",
    "linux     mips64le",
    "netbsd    386",
    "netbsd    amd64",
    "netbsd    arm",
    "openbsd   386",
    "openbsd   amd64",
    "openbsd   arm",
    "plan9     386",
    "plan9     amd64",
    "solaris   amd64",
    "windows   386",
    "windows   amd64"
];

mkdir 'builds';

foreach my $platform (@$os_arch) {
    my ( $os, $arch ) = split /\s+/, $platform;
    $ENV{'GOOS'}   = $os;
    $ENV{'GOARCH'} = $arch;
    system("go build -v -o builds/adjust_cli_${os}_${arch}");
}
