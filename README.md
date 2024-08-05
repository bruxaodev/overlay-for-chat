# **Chat Overlay**

This overlay application allows users to add a chat window from their stream and keep it always on top of other windows. Designed to enhance the streaming experience, this tool ensures that streamers can easily monitor and interact with their audience without switching between applications. The overlay is lightweight, easy to use, and perfect for maintaining engagement during live broadcasts.

## **Index**

### English
- [Features](#features)
- [Usage Steps](#usage-steps)
- [Development and Build Instructions](#development-and-build-instructions)

### Português
- [Recursos](#recursos)
- [Passos para Uso](#passos-para-uso)

---

## **English**

### **Features**
- **Always On Top:** Keeps the chat window overlaid on other windows.
- **User-Friendly:** Simple and intuitive interface for easy setup and use.
- **Customizable:** Adjust the size and position of the chat window to fit your needs.
- **Lightweight:** Minimal system resources required, ensuring smooth performance.
- **Transparency Support:** If your chat has transparency, it will work seamlessly with the overlay.

### **Usage Steps**
1. **Open the Program:** Run the program for the first time to generate the configuration file.
2. **Close the Program:** Once the `config.json` file is created, close the program.
3. **Edit the Configuration:** Open the `config.json` file and replace the default link with the link to your chat or any other content you want to overlay.
4. **Toggle Window Bar:** Press `Ctrl + Shift + X` to activate or deactivate the window bar and save the current state to the configuration.
5. **Customize Shortcuts:** You can also edit the keyboard shortcuts in the `config.json` file to suit your preferences.

### **Development and Build Instructions**
This project uses the [Wails3](https://v3alpha.wails.io/) framework for development.

#### **Prerequisites**
- Ensure you have [Wails3 installed](https://v3alpha.wails.io/getting-started/installation/).

#### **Run the Project**
1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Start the development server:
   ```sh
   wails dev
   ```

#### **Build the Project**
1. Build the project for production:
   ```sh
   wails build
   ```

---

## **Português**

### **Recursos**
- **Sempre no Topo:** Mantém a janela de chat sobreposta a outras janelas.
- **Fácil de Usar:** Interface simples e intuitiva para fácil configuração e uso.
- **Personalizável:** Ajuste o tamanho e a posição da janela de chat conforme suas necessidades.
- **Leve:** Requer recursos mínimos do sistema, garantindo um desempenho suave.
- **Suporte a Transparência:** Se o seu chat tiver transparência, ele funcionará perfeitamente com o overlay.

### **Passos para Uso**
1. **Abra o Programa:** Execute o programa pela primeira vez para gerar o arquivo de configuração.
2. **Feche o Programa:** Após a criação do arquivo `config.json`, feche o programa.
3. **Edite a Configuração:** Abra o arquivo `config.json` e substitua o link padrão pelo link do seu chat ou qualquer outro conteúdo que você deseja sobrepor.
4. **Ativar/Desativar Barra da Janela:** Pressione `Ctrl + Shift + X` para ativar ou desativar a barra da janela e salvar o estado atual na configuração.
5. **Personalize Atalhos:** Você também pode editar as teclas de atalho no arquivo `config.json` para se adequar às suas preferências.