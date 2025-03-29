<h1 align="center">Frieren (Discord Bot)</h1> 

![Frieren](assets/frieren.jpg)

<div align="center">
  <img src="https://img.shields.io/badge/Golang-gray?logo=go" height="30">
  <img src="https://img.shields.io/badge/Discord-gray?logo=discord" height="30">
  <img src="https://img.shields.io/badge/Youtube-gray?logo=youtube&logoColor=red" height="30">
  <img src="https://img.shields.io/badge/Docker-gray?logo=docker" height="30">
  <img src="https://img.shields.io/badge/PostgreSQL-gray?logo=postgresql" height="30">
  <img src="https://img.shields.io/badge/Redis-gray?logo=redis" height="30">
  <!-- <img src="https://img.shields.io/badge/License-MIT-green" height="30"> -->
</div>




<h2 align="center">ABOUT</h2> 

<div align="center">
  <img src="https://media.tenor.com/rHlkrXa_lW8AAAAM/anime-frieren.gif" align="right" height="230" width="250"/>
</div>

<div align="left">
This is a Discord bot written in Go, designed to bring fun and utility to your server! <br> While the bot's current functionality is simple, we're actively working on expanding its capabilities.

- 🎵 Music: Play, stop, skip, pause, and resume.
- 🎭 Roles by Reaction
- ⚙️ Manual Command Customization
- 🎮 Games in Frieren Style (Coming Soon)
- 📱 Gacha Game Support (Coming Soon)

The bot's mascot is <b>Frieren<b>.<br>
Built with [discordgo](https://github.com/bwmarrin/discordgo) and [disgolink](https://github.com/disgoorg/disgolink/tree/v3).


</div>

<br>
<hr>
<br>


### 🛠 Installation & Setup

Ensure you have the following installed:
- **Docker**
- **Git**

#### 1️⃣ Clone the Repository
```bash
git clone https://github.com/Snayt1k3/discord-bot
cd discord-bot
```

#### 2️⃣ Set Up Environment Variables

Create a `.env` file in the project root and add the following variables:
```env
DISCORD_TOKEN=""
BOT_STATUS=""
APPLICATION_ID=""
GUILD_ID=""
LAVALINK_ADDR=""
LAVALINK_PASS=""
LAVALINK_NODE_NAME=""
LavalinkSecure=""

SETTINGS_ADDRESS=""
SETTINGS_PORT=""

API_GATEWAY_ADDR=""
API_GATEWAY_PORT=""

SETTINGS_POSTGRES_HOST=""
SETTINGS_POSTGRES_PORT=""
SETTINGS_POSTGRES_USER=""
SETTINGS_POSTGRES_PASSWORD=""
SETTINGS_POSTGRES_DB=""
SETTINGS_POSTGRES_SSLMODE=""
SETTINGS_POSTGRES_TIMEZONE=""

REDIS_HOST=""
REDIS_PORT=""
REDIS_PASS=""
REDIS_DB=""
```

#### 3️⃣ Build and Run with Docker
```bash
docker-compose up --build -d
```

This will build and start the bot in a Docker container.


<br>
<hr>

🤝 Contributing

Feel free to submit issues, feature requests, or pull requests to improve Frieren!
