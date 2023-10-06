# nifi-tui

Unofficial terminal user interface (TUI) for [Apache NiFi](https://nifi.apache.org/)

## Implemented features

- [x] Connect to secure instance
- [x] View process group components
- [x] Create and delete components (processor, process group, etc.) in a process group
- [x] Create connections
- [x] Start/stop processors and ports
- [x] View and edit processor settings, properties and relationships

## Roadmap

There are many missing features compared to the web UI and there are no plans for this application to replace it.
The main motivation for this application is to provide an alternative method for developers and support people to view and configure NiFi when a browser is not available for any reason. So the current aim is to allow some basic configurations like:

- [ ] Allow configuring components other than processors
- [ ] Display the number of FlowFiles in a queue
- [ ] Allow viewing and configuring controller services, parameters, reporting tasks, etc.
- [ ] Add fuzzy search for components
- [ ] Open the log file where the exception displayed in the bulletin has occured
- [ ] Load password from file or env variable to make arguments more secure
- [ ] Refresh auth token when session has expired

It could, in theory, replace some toolkit functionalities with a nicer interface as well like:

- [ ] Export templates
- [ ] Collect and export status history

## Usage

To start _nifi-tui_, you must pass the url (defaults to [https://localhost:8443](https://localhost:8443)) where NiFi is running, user name and password as arguments

```bash
nifi-tui --url https://localhost:8443 --user-name "user-name" --password "password"
```

Refer to the help bar at the bottom of the application for the keyboard mappings in the current screen
