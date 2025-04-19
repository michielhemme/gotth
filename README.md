# GoTTH CLI

A command-line tool for bootstrapping and developing applications with the **GoTTH stack**:
- **Go** — the backend language
- **Templ** — HTML templating engine
- **TailwindCSS** — utility-first CSS framework
- **HTMX** — modern frontend interaction library

GoTTH wraps and manages essential CLI tools like `templ`, `air`, and `tailwindcss` to streamline development.

---

## 🚀 Features

- ⛔ Initialize a new GoTTH project with best practices
- ✅ Manage and build frontend assets using TailwindCSS
- ✅ Live reload support via `gotth air`
- ✅ Compile Templ templates
- ⛔ Simplified setup process

---

## 📦 Installation

Download the latest release from the [releases page](#) and place the binary in a folder that’s part of your system’s `PATH`.

### Adding to PATH

- **Windows**
  1. Press `Win + R`, type `sysdm.cpl`, and hit Enter.
  2. Under the **Advanced** tab, click **Environment Variables**.
  3. Under **System Variables**, find `Path` and click **Edit**.
  4. Add the folder where your GoTTH binary resides.

- **Linux / macOS**
  Add this line to your shell profile file (`~/.bashrc`, `~/.zshrc`, etc.):
  ```sh
  export PATH="$PATH:/path/to/your/gotth"
  ```
  Then reload your shell:
  ```sh
  source ~/.bashrc  # or ~/.zshrc
  ```

---

## 🛠️ Building from Source
To build GoTTH from source, clone the repository and ensure you have the necessary dependencies installed.
### Download Tools
You can download required CLI tools using:
```sh
go run ./tools/downloader
```
Or if you're using make:
```sh
make download-tools
```
### Build
After downloading the dependencies, compile the project:
```sh
go build -o ./tmp/gotth .
```
or if you're using make:
```sh
make build
```
The compiled binary will be located in the `./tmp` directory.

---

## 📄 License
This project is licensed under the MIT License.

---

## 🙌 Contributions
Pull requests and issues are welcome! If you'd like to contribute, feel free to open a PR or file a bug report.

---

## Feedback
Questions, feedback, or feature requests? Open an issue or reach out!