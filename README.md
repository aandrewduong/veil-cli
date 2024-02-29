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
3. **Enrollment**: Enroll in courses.
4. **Watch**: Watch the enrollment data for classes, notifying you if there is a waitlist or enrollment spot available.

## Prerequisites

- **Golang**: You need a version >=1.21.4 of [Go](https://go.dev/doc/install) installed.

## Configuration

For the tool to function correctly, "settings.csv" is required to be setup properly.

### settings.csv Parameters

| Parameter            | Description                                         | Example Values                               |
|----------------------|-----------------------------------------------------|----------------------------------------------|
| Username             | Your FHDA Username                                  | `00000000`                                   |
| Password             | Your FHDA Password                                  | `TestTestPassword123`                        |
| Term                 | Term                                                | `2024 Spring De Anza`                        |
| Subject              | Subject (Used for Class Search)                     | `MATH`                                       |
| Mode                 | Type of Task                                        | `Signup`                                     |
| CRNs                 | Course Refernce Numbers                             | `47520,44412,41846`                          |
| Webhook       | Discord Webhook URL (For notifications)                    | `https://discord.com/api/webhooks/[gone] `   |

To create a Discord Webhook, See [How to Create a Discord Webhook](https://hookdeck.com/webhooks/platforms/how-to-get-started-with-discord-webhooks).

To edit settings.csv, a spreadsheet editor is recommended. See [Rons Editor](https://www.ronsplace.ca/products/ronseditor) or Google Sheets.

## Compilation

To compile this program, run build.sh for the program to be compiled.

## Usage

To run this program, run
```
go run .
```

## Mode

- **Signup**: Enroll in classes with specified CRNs (Course Refence Numbers).
- **Search**: Search for all the sections based on given term, section and subject.
- **Transcript**: Export your "unofficial transcript", or data of previously enrolled courses.
- **Watch**: Monitor enrollment data, sending a notification if a waitlist or enrollment spot is available.

## Notification Example

![Notification](https://media.discordapp.net/attachments/1022240002408730644/1212615611465859082/image.png?ex=65f27b4b&is=65e0064b&hm=75468e9840762051800341e47605d339dbe3c50f80e45e6a678131d099eebb43&=&format=webp&quality=lossless)
![Notification](https://media.discordapp.net/attachments/1022240002408730644/1212616160810504212/image.png?ex=65f27bce&is=65e006ce&hm=6a8b307714e536217d06d7364351138b5b09171cc5d2ccb7e70e669bcec83e10&=&format=webp&quality=lossless)


## Version 2 vs Version 1

- **Code Refractor**: Major code refractor in this version.
- **Speed**: Slight speed improvements.
- **Watch**: Brand new watch feature, useful for sniping waitlist or enrollment spots.