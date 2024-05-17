# Veil-V2

Veil is an open-source program written in Golang designed to efficiently scrape, process, and manage class and college enrollment data at De Anza and Foothill College. In addition to offering a seamless way to search and export class data, it also supports class enrollment and class enrollment monitoring.

## About Me

I am a Junior at De Anza College. I am transferring in Fall 2024, so there's no point of keeping this code with me any longer.
[Linkedin](https://www.linkedin.com/in/andrew-duong-3a9931259/)
Hope people can build off this project, enjoy!

## Table of Contents

- [Key Features](#key-features)
- [Prerequisites](#prerequisites)
- [Configuration](#configuration)
- [Compilation](#compilation)
- [Usage](#usage)
- [Modes](#modes)
- [Notification Example](#notification-example)


## Key Features

1. **Class Search & Export**: Ability to search for classes and export the results in CSV format.
2. **Unofficial Transcript**: Retrieve and export your previously enrolled courses in CSV format.
3. **Enrollment**: Enroll in courses in lighting speed.
4. **Watch**: Watch the enrollment data for classes, notifying you if there is a waitlist or enrollment spot available and signs up for the waitlist spot for you.

## Prerequisites

- **Golang**: You need a version >=1.21.4 of [Go](https://go.dev/doc/install) installed.

## Configuration

For the tool to function correctly, "settings.csv" is required to be setup properly.

### settings.csv Parameters

| Parameter            | Description                                         | Example Values                               |
|----------------------|-----------------------------------------------------|----------------------------------------------|
| Username             | Your FHDA Username                                  | `00000000`                                   |
| Password             | Your FHDA Password                                  | `TestTestPassword123`                        |
| Term                 | Term                                                | `2024 Summer De Anza`                        |
| Subject              | Subject (Used for Class Search)                     | `MATH`                                       |
| Mode                 | Type of Task                                        | `Signup`                                     |
| CRNs                 | Course Reference Numbers                             | `47520,44412,41846`                          |
| Webhook              | Discord Webhook URL (For Notifications)             | `https://discord.com/api/webhooks/[gone] `   |
| SavedRegistrationTime | Saved Registration Time (Use Signup Mode)          | (Do not Edit on your own)                    |

To create a Discord Webhook, See [How to Create a Discord Webhook](https://hookdeck.com/webhooks/platforms/how-to-get-started-with-discord-webhooks).

To edit settings.csv, a spreadsheet editor is recommended. See [Rons Editor](https://www.ronsplace.ca/products/ronseditor) or Google Sheets.

## Compilation

To compile this program, run build.sh for the program to be compiled.

## Usage

To run this program, run
```
go run .
```

## Example Scenarios

Scenario 1: 
If I want the program to signup classes for me, I would run Signup mode and fill in settings.csv. If I want complete automation - 
run Signup or Release mode first because it'll save your registration time. After the registration time is saved in settings.csv, the program will sleep until 5 minutes before registration opens.

Scenario 2:
I'm currently looking at a class but the waitlist is full! I would run Watch mode and fill in settings.csv. If you see "Getting Enrollment Data (XXXXX)" this is completely normal. Once a waitlist spot is open, Watch mode will start a
Signup task and try to enroll you in that waitlist spot.

Scenario 3:
How do I get the catalog of a unpublished term? You would still use Search mode. The task will build the Term ID locally instead of using FHDA's API because that term isn't published yet.

## Mode

- **Release**: Essentially the same as Signup mode. However, the program will wait until (SavedRegistrationTime - 5 minutes) before running. For example, if my registration time was 8:00 AM, the program will sleep until 7:55 AM. (Meaning you can run this program overnight!)
- **Signup**: Enroll in classes with specified CRNs (Course Refence Numbers).
- **Search**: Search for all the sections based on given term, section and subject.
- **Transcript**: Export your "unofficial transcript", or data of previously enrolled courses.
- **Watch**: Monitor enrollment data, sending a notification if a waitlist or enrollment spot is available. If a waitlist spot is available, it'll run a Signup mode task and enroll in the waitlist.

![image](https://github.com/aandrewduong/veil-v2/assets/135930507/e6e862df-2fde-4015-9095-d9e4818047f3)

Encounter any bugs or have questions/feature requests? Message me on Discord! @ascendingz
