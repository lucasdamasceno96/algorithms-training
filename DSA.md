# 📘 DSA Cheat Sheet – Fundamentos Antes do Primeiro Código

> Objetivo: Este documento é seu mapa mental inicial para começar Data Structures & Algorithms (DSA) com mentalidade de engenheiro de plataforma/cloud.

---

# 1️⃣ O que é DSA?

**Data Structures (Estruturas de Dados)** = Como organizamos dados na memória.

**Algorithms (Algoritmos)** = Como processamos esses dados de forma eficiente.

Em engenharia de plataforma:

- Estrutura errada = mais memória, mais latência, mais custo.
- Algoritmo ruim = CPU alta, autoscaling desnecessário.

DSA é sobre:

- Performance
- Escalabilidade
- Uso eficiente de memória

---

# 2️⃣ Complexidade de Tempo e Espaço

## ⏱ Complexidade de Tempo (Time Complexity)

Quanto o tempo cresce conforme o input cresce.

## 💾 Complexidade de Espaço (Space Complexity)

Quanto de memória extra o algoritmo consome.

---

# 📊 Tabela de Big-O

| Notação    | Nome        | Crescimento              | Exemplo típico          |
| ---------- | ----------- | ------------------------ | ----------------------- |
| O(1)       | Constante   | Não cresce               | Acessar índice de array |
| O(log n)   | Logarítmico | Cresce lentamente        | Binary Search           |
| O(n)       | Linear      | Cresce proporcionalmente | Loop simples            |
| O(n log n) | Linear-log  | Muito comum em sorting   | Merge Sort              |
| O(n²)      | Quadrático  | Cresce muito rápido      | Dois loops aninhados    |
| O(2^n)     | Exponencial | Explode rapidamente      | Backtracking            |
| O(n!)      | Fatorial    | Impraticável             | Permutações             |

---

# 3️⃣ Estruturas de Dados Fundamentais

---

## 📦 Array (Python List / Go Array)

### O que é?

Bloco contínuo de memória.

### Características:

- Acesso por índice: O(1)
- Inserção no final: O(1) amortizado
- Inserção no meio: O(n)

### Analogia Infra:

Como um bloco contíguo em memória RAM. Cache-friendly.

### Python

```python
numbers: list[int] = [1, 2, 3]
value: int = numbers[0]
```

### Go

```go
arr := [3]int{1, 2, 3}
value := arr[0]
```

---

## 🧩 Slice (Go)

### O que é?

Abstração dinâmica sobre array.

### Características:

- Pode crescer
- Internamente aponta para array

### Estrutura interna:

- Pointer
- Length
- Capacity

### Go

```go
nums := []int{1, 2, 3}
nums = append(nums, 4)
```

⚠ Se ultrapassar capacity → novo array é criado (realocação).

---

## 🗂 HashMap (Python Dict / Go Map)

### O que é?

Tabela de espalhamento baseada em hash.

### Complexidade:

- Busca: O(1) médio
- Inserção: O(1) médio
- Pior caso: O(n)

### Analogia Infra:

Como um DNS: você resolve chave → valor rapidamente.

### Python

```python
users: dict[str, int] = {"alice": 1}
id: int = users["alice"]
```

### Go

```go
users := make(map[string]int)
users["alice"] = 1
id := users["alice"]
```

---

## 📚 Stack (Pilha)

### Regra:

LIFO (Last In First Out)

### Uso:

- Recursão
- Undo
- Call stack

### Python

```python
stack: list[int] = []
stack.append(1)
value: int = stack.pop()
```

---

## 📬 Queue (Fila)

### Regra:

FIFO (First In First Out)

### Analogia Infra:

Kafka, RabbitMQ → consumidores processam em ordem.

### Python

```python
from collections import deque

queue: deque[int] = deque()
queue.append(1)
value: int = queue.popleft()
```

---

# 4️⃣ Algoritmos Fundamentais

---

## 🔎 Binary Search

Requer array ordenado.

Complexidade: O(log n)

Dividimos o problema pela metade a cada passo.

---

## 🔁 Sorting Principais

| Algoritmo   | Complexidade     | Estável? |
| ----------- | ---------------- | -------- |
| Bubble Sort | O(n²)            | Sim      |
| Merge Sort  | O(n log n)       | Sim      |
| Quick Sort  | O(n log n) médio | Não      |
| Heap Sort   | O(n log n)       | Não      |

---

# 5️⃣ Conceitos Fundamentais Antes do Primeiro Código

---

## 🧠 Recursão

Função que chama a si mesma.

Elementos obrigatórios:

- Base case
- Recursive case

---

## 🧮 Iteração vs Recursão

Recursão usa stack.
Iteração usa loop explícito.

Infra analogy:
Recursão mal controlada = stack overflow = crash.

---

## 📐 Two Pointers

Usado para:

- Arrays ordenados
- Sliding window

Complexidade típica: O(n)

---

## 🪟 Sliding Window

Mantém janela dinâmica sobre array.

Muito usado para:

- Substrings
- Métricas de janela (observabilidade, rate limiting)

---

## 🌳 Noções Futuras (Próximo Nível)

- Linked List
- Binary Tree
- Binary Search Tree
- Heap
- Graph
- DFS / BFS
- Dynamic Programming

---

# 6️⃣ Diferença Crítica: Array vs HashMap

| Critério       | Array      | HashMap     |
| -------------- | ---------- | ----------- |
| Ordem          | Mantém     | Não garante |
| Acesso         | Por índice | Por chave   |
| Busca          | O(n)       | O(1) médio  |
| Memória        | Contígua   | Espalhada   |
| Cache friendly | Sim        | Não tanto   |

---

# 7️⃣ Modelo Mental Para Entrevistas Big Tech

Sempre pergunte:

1. Qual é o tamanho do input (n)?
2. Preciso de busca rápida?
3. Preciso manter ordem?
4. O dataset cabe em memória?
5. Posso pré-processar?

---

# 8️⃣ Erros Clássicos de Iniciantes

- Não calcular complexidade
- Ignorar edge cases
- Não considerar memória
- Usar estrutura errada
- Resolver força bruta primeiro

---

# 9️⃣ Setup Mental Antes de Resolver Qualquer Problema

1. Entenda o input
2. Entenda o output
3. Liste edge cases
4. Defina estrutura de dados
5. Defina complexidade alvo
6. Só então comece a codar

---

# 🔟 Roadmap Inicial de Estudo

Semana 1:

- Big-O
- Arrays
- HashMap
- Two Pointers

Semana 2:

- Stack
- Queue
- Recursão

Semana 3:

- Linked List
- Binary Search
- Sorting

Semana 4:

- Trees
- BFS/DFS

---

# 🏁 Conclusão

DSA não é sobre decorar.
É sobre:

- Modelar problema
- Escolher estrutura correta
- Controlar custo computacional

Para engenheiros de plataforma:

Algoritmo ruim = mais instâncias no cluster.
Estrutura errada = mais consumo de memória.

Eficiência = menos custo em cloud.

---

Fim do Cheat Sheet 🚀
