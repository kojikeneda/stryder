# Stryder

Stryder is a powerful server framework designed to collect data on internet service providers. It is a part of a larger website that helps people understand vendor reliability and can be used for AIOps data enrichment use cases. 

Stryder collects data in various forms such as direct tests against a provider's service or information found on their site such as a status page. It can be deployed as an on-box agent or as a third-party poller.

## Architecture

Stryder consists of the following key components:

1. **Stryder Core**: The central orchestrator of the system. It schedules and manages data collection, determines where data should be stored, and processes requests to schedule the execution of data collection.

2. **Stryder Plugins**: Modules that provide integrations for data collection and storage. They interact with the Stryder Core to perform these tasks.

3. **Stryder Agent**: A packaged combination of Stryder Core and Stryder Plugins. The agent can be deployed on servers either as an on-box agent or as a third-party poller.

Additionally, Stryder monitors its own execution metrics, including CPU usage, memory usage, network traffic, and disk usage. It reports these metrics to various observability platforms.

## Project Structure

The Stryder project is organized as follows:


Each directory has a specific purpose:

- `cmd/`: Contains the application's main files. Each file in the `cmd/` directory corresponds to a buildable artifact.

- `pkg/`: Contains the majority of the code. It's where other, non-`main` packages are defined.

- `api/`: Definitions of API types and protocol definition files.

- `scripts/`: Scripts to perform various build, install, analysis, etc operations.

- `test/`: Additional external test apps and test data.

## Getting Started

This section will be updated with instructions on how to install and run Stryder.

## Contributing

This section will be updated with instructions on how to contribute to the Stryder project.

## License

This section will be updated with information on the project's license.
