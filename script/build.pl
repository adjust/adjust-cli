#!/usr/bin/perl

use strict;
use warnings;
use 5.018;

use JSON;
use LWP::UserAgent;
use autodie;
use Cwd;

my $os_arch = [
    "darwin    386",
    "darwin    amd64",
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

system('goem bundle q');
$ENV{'GOPATH'} = getcwd . '/.go';

my $json;

my $ua   = LWP::UserAgent->new;
my $resp = $ua->get('https://api.github.com/repos/adjust/adjust-cli/tags');

if ( $resp->is_success ) {
    $json = decode_json( $resp->decoded_content );
} else {
    die $resp->status_line;
}

my $build_base   = 'builds';
my $build_dir    = "$build_base/$json->[0]->{name}/";
my $build_latest = "$build_base/latest";

mkdir $build_base unless -d $build_base;

if ( -d $build_dir ) {
    say "nothing to do";
} else {
    mkdir $build_dir;
    eval {
        foreach my $platform (@$os_arch) {
            my ( $os, $arch ) = split /\s+/, $platform;
            $ENV{'GOOS'}   = $os;
            $ENV{'GOARCH'} = $arch;
            say "Building for $os / $arch";
            eval {
                system("go build -o $build_dir/adjust_cli_${os}_${arch}") == 0 or die $!; };
            say "build failed for $os/$arch: $@" if $@;
            say "done building $os / $arch";
        }
    };
}

eval {
    symlink $build_dir, $build_latest;
};
