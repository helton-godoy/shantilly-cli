# PRD - Shantilly-CLI TUI

## Product Requirements Document

### Visão do Produto
Criar uma interface terminal moderna (TUI) para shell scripts que substitua ferramentas tradicionais como `dialog` e `whiptail`.

### Problema a Resolver
- Ferramentas atuais (dialog/whiptail) são visualmente datadas
- Falta de consistência cross-platform
- Experiência do usuário limitada
- Dificuldade de personalização

### Solução Proposta
Shantilly-CLI: TUI moderna baseada em Charmbracelet com:
- Interface visualmente atrativa
- Cross-platform (Linux, macOS, Windows)
- Fácil integração com shell scripts
- Extensível e personalizável

### Requisitos Funcionais

#### 1. Dialog Boxes
- Menus interativos com múltiplas opções
- Navegação por teclado
- Suporte a atalhos
- Seleção única ou múltipla

#### 2. Forms
- Campos de diversos tipos (text, email, number, password)
- Validação em tempo real
- Máscaras e formatação
- Obrigatoriedade de campos

#### 3. Progress Bars
- Indicadores visuais de progresso
- Múltiplos estilos disponíveis
- Porcentagem e tempo estimado
- Cancelamento opcional

#### 4. File Selectors
- Navegação de diretórios
- Filtros por extensão
- Seleção única ou múltipla
- Preview de arquivos

#### 5. Confirmations
- Diálogos sim/não/Cancelar
- Mensagens personalizáveis
- Timeout automático
- Valor padrão configurável

### Requisitos Técnicos

#### Stack Tecnológico
- **Linguagem:** Golang 1.21+
- **TUI Framework:** Charmbracelet (bubbletea, lipgloss, bubbles)
- **Build:** Makefile para multi-plataforma
- **Distribuição:** Binário único

#### Arquitetura
```
cmd/shantilly/main.go          # Entry point CLI
pkg/tui/                       # Componentes TUI
├── dialog.go                  # Dialog boxes
├── form.go                    # Formulários
├── progress.go                # Progress bars
└── selector.go                # File selectors
pkg/config/config.go           # Configurações
```

#### Comandos CLI
```bash
shantilly dialog --title "Menu" --options "Op1,Op2,Op3"
shantilly form --fields "name,email,age"
shantilly progress --steps "Build,Test,Deploy"
shantilly select --path "/home" --filter "*.go"
shantilly confirm --message "Continue?"
```

### Critérios de Sucesso

#### Funcionalidade
- [ ] Todos os componentes TUI funcionando
- [ ] Cross-platform compatibility
- [ ] Integração com shell scripts
- [ ] Performance adequada (<100ms startup)

#### Qualidade
- [ ] 90%+ test coverage
- [ ] Zero vulnerabilidades de segurança
- [ ] Documentação completa
- [ ] Binário <10MB

#### Experiência do Usuário
- [ ] Interface intuitiva
- [ ] Navegação por teclado completa
- [ ] Tratamento de erros graceful
- [ ] Feedback visual claro

### User Stories

#### Como desenvolvedor shell script
Eu quero criar menus interativos bonitos
Para que meus scripts tenham melhor UX

#### Como sysadmin
Eu quero formulários de coleta de dados
Para automatizar tarefas com entrada de usuário

#### Como DevOps
Eu quero barras de progresso visuais
Para monitorar processos longos

#### Como usuário final
Eu quero interfaces consistentes
Para ter experiência familiar em diferentes scripts

### MVP (Minimum Viable Product)

#### Funcionalidades Essenciais
1. Dialog boxes básicas
2. Forms simples
3. Progress bars
4. Cross-platform básico

#### Funcionalidades Futuras (v1.1)
1. File selectors avançados
2. Temas e personalização
3. Plugins e extensões
4. Web TUI integration

### Riscos e Mitigações

#### Riscos Técnicos
- **Complexidade Charmbracelet:** Mitigar com protótipos rápidos
- **Performance Cross-platform:** Testar em todos os alvos
- **Integração Shell:** Testar com diversos shells

#### Riscos de Mercado
- **Adoção:** Mitigar com documentação e exemplos
- **Concorrência:** Diferenciar com UX superior

### Timeline Estimada

#### Sprint 1 (2 semanas)
- Setup do projeto Go
- Implementação dialog boxes
- Forms básicos

#### Sprint 2 (2 semanas)  
- Progress bars
- Cross-platform testing
- Documentação inicial

#### Sprint 3 (1 semana)
- File selectors
- Confirmations
- Release v1.0.0

### Métricas de Sucesso

#### Técnicas
- Performance: <100ms startup
- Memory: <50MB usage
- Size: <10MB binary

#### Adoção
- Downloads: >1000/mês
- Stars: >100 no GitHub
- Issues: <5 bugs críticos

---

**Status:** ✅ **PRD COMPLETO**  
**Next:** Arquitetura técnica  
**Responsível:** Architect Agent  
**Timeline:** 5 semanas total
