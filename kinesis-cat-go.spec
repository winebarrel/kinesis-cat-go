Name:		kinesis-cat-go
Version:	0.1.3
Release:	1%{?dist}
Summary:	Amazon Kinesis cli for put JSON data

Group:		Development/Tools
License:	MIT
URL:		https://github.com/winebarrel/kinesis-cat-go
Source0:	kinesis-cat-go-%{version}.tar.gz
BuildRoot:	%(mktemp -ud %{_tmppath}/%{name}-%{version}-%{release}-XXXXXX)

%description
Amazon Kinesis cli for put JSON data

%prep
%setup -q

%build
make

%install
rm -rf %{buildroot}
mkdir -p %{buildroot}/usr/bin
make install DESTDIR=%{buildroot} PREFIX=/usr

%clean
rm -rf %{buildroot}

%files
%defattr(-,root,root,-)
/usr/bin/kinesis-cat
