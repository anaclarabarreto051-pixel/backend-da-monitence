const form = document.querySelector("form");

form.addEventListener("submit", async (event) => {
  event.preventDefault();

  const payload = {
    name: document.getElementById("name").value,
    email: document.getElementById("email").value,
    password: document.getElementById("password").value,
    phone: document.getElementById("phone").value,
  };

  try {
    const response = await fetch("http://localhost:8081/api/users", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });

    const text = await response.text(); // pega resposta como texto primeiro
    let data = {};
    try { data = text ? JSON.parse(text) : {}; } catch {}

    if (!response.ok) {
      console.error("Erro HTTP:", response.status, data || text);
      alert(`Erro ao cadastrar (HTTP ${response.status})`);
      return;
    }

    console.log("Sucesso:", data);
    alert("Usuário cadastrado com sucesso!");
    form.reset();
  } catch (error) {
    console.error("Erro de rede:", error);
    alert("Falha de rede. O backend tá ligado?");
  }
});