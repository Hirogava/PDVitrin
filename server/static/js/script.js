document.querySelectorAll('.filter-btn').forEach(button => {

    button.addEventListener('click', () => {
      button.classList.toggle('active');
    });
  });