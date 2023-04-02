document.addEventListener("DOMContentLoaded", () => {
  const actionBtn = document.getElementById("actionBtn");

  if (actionBtn) {
    actionBtn.addEventListener("click", () => {
      fetch('http://localhost:8080/api/hello')
        .then(response => response.json())
        .then(data => alert(data.message))
        .catch(error => console.error(error))
    });
  }
});