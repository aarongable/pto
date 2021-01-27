# PTO Tracker

A simple tool to help you compute how many days of Paid Time Off you have
accrued as a result of your work. Particularly helpful for employees of
companies which use "unlimited time off", in order to track things themselves
and ensure they take reasonable amounts of time off.

## Installation

```shell
go get github.com/aarongable/pto
```

## Usage

If you have been working since the unix epoch, and earn 0.125 hours of PTO per
hour that you work:

```shell
pto -startDate 1970-01-01 -hph 0.125
```

If you have been working since Y2K, and earn 25 days of PTO per year:

```shell
pto -startDate 2000-01-01 -dpy 25
```
