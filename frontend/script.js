const form = document.getElementById('registrationForm');

form.addEventListener('submit', function(event) {
    event.preventDefault();

    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const phone = document.getElementById('phone').value;

    if (!name || !email || !password || !phone) {
        alert('Please fill in all fields');
        return;
    }

const payload = {
  name: document.getElementById('name').value,
  email: document.getElementById('email').value,
    password: document.getElementById('password').value,
   phone: document.getElementById('phone').value
};

console.log(payload);

  //  event.preventDefault();
    //const name = document.getElementById('name').value;
   // const email = document.getElementById('email').value;
  //  const password = document.getElementById('password').value;
   // const phone = document.getElementById('phone').value;
   // console.log({ name, email, password, phone });

fetch('http://localhost:8080/api/users', {   
  method: 'POST',
    headers: {
       'Content-Type': 'application/json'
    },
    body: JSON.stringify(payload)
})
.then((response) => {
    console.log('Status:', response.status);
        return response.json();
})
.then(data => {
    console.log('Success:', data); 
    alert('Usuário cadastrado com sucesso!');
    form.reset();
})
.catch(error => {
    console.error('Error:', error);
    alert('Erro ao cadastrar usuário. Por favor, tente novamente.');
});
});
