# Pomodoro CLI

## Overview

The Pomodoro CLI is a simple command-line tool for implementing the Pomodoro Technique—a time management method.

## Features

Pomodoro Timer: Set a timer for your work sessions, and the CLI will notify you when the time is up.

Customizable Settings: Adjust the duration of Pomodoro and break sessions according to your preferences.

## Installation

Download the binary(windows) or clone this repository and build it. Make sure you have Go instaled on the most recent version.

## Start a Pomodoro Session:

```bash
pomo
```

This will initiate a 25-minute Pomodoro session.

## Adjust Timer Duration:

To set a custom duration for the Pomodoro session or break, use the -min flag:

```bash
pomo --min=5
```

This starts a Pomodoro session with a duration of 5 minutes.

## View Help:
For more options and information, use the --help flag:

```bash
pomo --help
```
