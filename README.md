<h1 align="center">Frieren (Discord Bot)</h1> 

![Frieren](assets/frieren.jpeg)

<div align="center">
    <img src="https://img.shields.io/badge/Golang-gray?logo=go" height="30">
    <img src="https://img.shields.io/badge/Discord-gray?logo=discord" height="30">
    <img src="https://img.shields.io/badge/Youtube-gray?logo=youtube&logoColor=red" height="30">
    
</div>




<h2 align="center">ABOUT</h2> 

<div align="center">
  <img src="assets/frieren_1.png" align="right" height="240" width="250"/>
</div>

<div align="left">
This is a Discord bot written in Go, made to bring fun and utility to your server!  
The current features are simple, but we’re actively adding more.

<!-- - 🎵 **Music**: play, pause, skip, and stop tracks. -->
- 🎭 **Reaction Roles**: assign roles to users via message reactions.
- 🛡️ **Automode**: banned-word filter, anti-caps, and anti-link protection.
- 🎮 **Frieren-Style Games (Coming Soon)**: calm and relaxing mini-games inspired by Frieren.

The bot’s mascot is **Frieren**.<br>
Built with [discordgo](https://github.com/bwmarrin/discordgo) and [disgolink](https://github.com/disgoorg/disgolink/tree/v3).
</div>

<br>
<br>
<br>
<br>

<h2 align="center">DEVELOPMENT</h2>

<div align="center">
  <img src="assets/frieren_2.png" align="right" height="250" width="250"/>
</div>

<div align="left">

**Setup**  
1. Run `make grpc-init` to generate gRPC files.  
2. Create a `.env` file based on `.env_example`.  

**Optional (Swagger)**  
- Run `make docs` to generate API docs.  
- Open [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) to view them.

**Run the Bot**  
```bash
make run
```

</div>
<hr>


