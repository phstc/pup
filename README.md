# pup

Start, stop and do not monitor processes.

## Usage

### Configuration

Create the `pup.json` configuration file in your home directory: `~/pup.json`.

```json
{
  "apps": [
  {
    "name": "railsapp",
    "cmd": "bundle exec rails s -p4000",
    "chdir": "/Users/pablo/workspace/railsapp",
    "pidfile": "/Users/pablo/railsapp.pid"
  },
  {
    "name": "standaloneapp",
    "cmd": "bin/start",
    "chdir": "/Users/pablo/workspace/standaloneapp",
    "pidfile": "/Users/pablo/standaloneapp.pid"
  },
  {
    "name": "sinatraapp",
    "cmd": "rackup",
    "chdir": "/Users/pablo/workspace/sinatraapp",
    "pidfile": "/Users/pablo/sinatraapp.pid"
  }
  ]
}
```

### Available commands:

```bash
Usage:
        pup (start|stop|restart|status) [<app-name>]

```

```bash
$ pup status
railsapp no such process
standaloneapp no such process
sinatraapp no such process

$ pup start railsapp
starting railsapp with pid 10268

$ pup status railsapp
railsapp is running with pid 10268

$ pup stop railsapp
stopping railsapp with pid 10268

$ pup status railsapp
railsapp no such process
```
