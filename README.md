<h1 align="center">✨ Frieren ✨</h1>
<p align="center"><i>A calm, friendly Discord bot written in Go.</i></p>

<p align="center">
  <img src="assets/frieren.jpeg" alt="Frieren" width="100%">
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white" height="28" alt="Go">
  <img src="https://img.shields.io/badge/Discord-5865F2?logo=discord&logoColor=white" height="28" alt="Discord">
  <img src="https://img.shields.io/badge/PRs-welcome-brightgreen" height="28" alt="PRs Welcome">
</p>

<p align="center">
  <a href="#-about">About</a> •
  <a href="#-features">Features</a> •
  <a href="#-development">Development</a> •
  <a href="#-contributing">Contributing</a>
</p>


## 🌸 About

<div align="center">
  <img src="assets/frieren_1.png" align="right" height="240" width="250" alt="Frieren mascot"/>
</div>

This is a **Discord bot written in Go**, made to bring fun and utility to your server!
The current features are simple, but we're actively adding more.

The bot's mascot is **Frieren**.
Built with [discordgo](https://github.com/bwmarrin/discordgo) and [disgolink](https://github.com/disgoorg/disgolink/tree/v3).


### 🎁 Features

- 🎭 **Reaction Roles** — assign roles to users via message reactions.
- 🛡️ **Automode** — banned-word filter, anti-caps, and anti-link protection.
- 🎮 **Frieren-Style Games** *(coming soon)* — calm and relaxing mini-games inspired by Frieren.
- 🎵 **Music** *(planned)* — play, pause, skip, and stop tracks.


## 🛠️ Development

<div align="center">
  <img src="assets/frieren_2.png" align="right" height="250" width="250" alt="Frieren development"/>
</div>

### Setup

1. Run `make grpc-init` to generate gRPC files.
2. Create a `.env` file based on `.env_example`.

### Optional (Swagger)

- Run `make docs` to generate API docs.
- Open [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) to view them.

### Run the Bot

```bash
make run
```

<br clear="right">


## 🤝 Contributing

Feel free to submit issues, feature requests, or pull requests to improve Frieren!
