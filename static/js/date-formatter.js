document.addEventListener("DOMContentLoaded", function () {
  // Seleciona todos os elementos que possuem um timestamp
  const timestamps = document.querySelectorAll(".timestamp");

  timestamps.forEach((el) => {
      const rawTimestamp = el.getAttribute("data-timestamp"); // Obtém o timestamp original

      if (rawTimestamp) {
          // Converte para um objeto Date
          const date = new Date(rawTimestamp);

          // Formata para o padrão brasileiro (dd/mm/aaaa hh:mm)
          const formattedDate = date.toLocaleString("pt-BR", {
              day: "2-digit",
              month: "2-digit",
              year: "numeric",
              hour: "2-digit",
              minute: "2-digit"
          });

          // Insere o timestamp formatado na célula da tabela
          el.textContent = formattedDate;
      }
  });
});
