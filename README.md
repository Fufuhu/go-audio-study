

# 動作確認環境

+ Ubuntu 20.04
+ macOS Monterey 12.6.3

# 前提ライブラリ等の導入

```bash
sudo apt install portaudio19-dev
```

macOS Monterey 12.6.3の場合は
```bash
brew install portdudio
```

https://github.com/hegedustibor/htgo-tts の動作に、
https://github.com/hajimehoshi/oto/ を使っているので以下を実行

```bash
sudo apt install libasound2-dev
```