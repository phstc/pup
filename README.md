# pup

Start, stop and do not monitor processes in development.

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
