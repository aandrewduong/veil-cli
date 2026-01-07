# Veil-CLI

Veil is a command-line interface written in Go designed for the automated management of course enrollment at De Anza and Foothill College. It streamlines class searching, data exporting, and enrollment monitoring by interfacing directly with the college's registration systems.

This version is optimized for the latest MyPortal and StudentRegistrationSsb updates.

---

## Table of Contents

- [Key Features](#key-features)
- [Prerequisites](#prerequisites)
- [Installation and Setup](#installation-and-setup)
- [Configuration](#configuration)
- [Compilation](#compilation)
- [Usage](#usage)
- [Operation Modes](#operation-modes)
- [Example Scenarios](#example-scenarios)

---

## Key Features

- **Course Search and Export**: Query available courses and export results to CSV format.
- **Academic Transcript**: Retrieve and export comprehensive unofficial transcript data.
- **Automated Enrollment**: Execute course registration with high efficiency.
- **Enrollment Monitoring**: Track course availability with real-time Discord notifications and automated waitlist enrollment.
- **Multi-Account Support**: Manage multiple student accounts and tasks simultaneously via a concurrent task engine.

---

## Prerequisites

- **Go (version 1.22.0 or higher)**: Installation instructions are available at the [official Go website](https://go.dev/doc/install).

---

## Installation and Setup

### 1. Clone the Repository
Clone the Veil repository to your local directory:

```sh
git clone https://github.com/aandrewduong/veil-cli.git
cd veil-cli
```

### 2. Install Dependencies
Initialize and download the required Go modules:

```sh
go mod tidy
```

### 3. Configure the Environment
Create a `config.json` file in the root directory to define your accounts and tasks. Refer to the [Configuration](#configuration) section for details.

---

## Configuration

Veil utilizes a `config.json` file to manage task parameters and credentials.

### Configuration Schema

```json
{
  "tasks": [
    {
      "username": "00000000",
      "password": "SecurePassword123",
      "term": "2025 Winter De Anza",
      "subject": "MATH",
      "mode": "Signup",
      "crns": ["47520", "44412"],
      "webhook_url": "https://discord.com/api/webhooks/...",
      "registration_time": "12/24/2025 08:00 AM"
    }
  ]
}
```

### Field Definitions

| Parameter | Description | Example |
|-----------|-------------|---------|
| `username` | FHDA Student ID | `00000000` |
| `password` | FHDA Account Password | `SecurePassword123` |
| `term` | Target academic term | `2025 Winter De Anza` |
| `subject` | Subject code for search | `MATH` |
| `mode` | Operation mode (e.g., `Signup`, `Watch`, `Classes`) | `Signup` |
| `crns` | Array of Course Reference Numbers | `["47520", "44412"]` |
| `webhook_url` | Discord Webhook URL for notifications | `https://discord.com/...` |
| `registration_time` | Registration start time (required for `Release` mode) | `01/02/2006 08:00 AM` |

---

## Compilation

To compile the source code into a binary executable:

```sh
make build
```

Alternatively, using the provided build script:

```sh
bash build.sh
```

---

## Usage

Execute the program directly using the Makefile:

```sh
make run
```

Or using the Go toolchain:

```sh
go run main.go
```

Alternatively, run the compiled binary:

```sh
./veil-cli
```

---

## Operation Modes

| Mode | Description |
|------|-------------|
| **Release** | Scheduled execution. The task idles until 5 minutes prior to the `registration_time`, maintaining an active session for immediate enrollment. |
| **Signup** | Immediate enrollment. Attempts to register for the provided CRNs as soon as the task starts. |
| **Classes** | Course discovery. Scrapes all available sections for the specified term and subject, exporting them to a CSV file. |
| **Transcript** | Data archival. Retrieves the student's unofficial transcript and exports it to a CSV file. |
| **Watch** | Enrollment tracking. Continuously monitors course availability. If a spot becomes available, it notifies the user via Discord and attempts auto-enrollment. |

---

## Example Scenarios

### Scenario 1: Automated Registration Day Execution
To ensure enrollment at the exact moment registration opens:
1. Configure a task with the `Release` mode and specify your exact `registration_time`.
2. Initialize Veil. The engine will maintain your session and begin the signup sequence exactly 5 minutes before your registration window opens.

### Scenario 2: Monitoring Full Waitlists
To secure a spot in a class that is currently full:
1. Configure a task with the `Watch` mode and include the target CRNs.
2. Veil will monitor the enrollment status. When a vacancy occurs, the program will trigger a `Signup` sequence and alert you via Discord.

### Scenario 3: Accessing Early Course Data
To view the course catalog for a future term before it is officially published in the portal:
1. Execute the `Classes` mode for the target term.
2. Veil generates the necessary term identifiers locally, allowing it to retrieve section data regardless of portal visibility.

---


<img width="1644" height="890" alt="image" src="https://github.com/user-attachments/assets/d1128495-8a71-4892-889c-c4771cf3a6ce" />
