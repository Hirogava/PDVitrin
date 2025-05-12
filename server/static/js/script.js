  const selectedFilters = {
    specialization: new Set(),
    theme: new Set(),
  };

  let debounceTimeout;
const pagination = document.getElementById('pagination');
function sendFiltersToAPI() {
  const jsonData = {
    specializations: { id: Array.from(selectedFilters.specialization) },
    thematics: { id: Array.from(selectedFilters.theme) },
  };

  fetch('/api/filter-projects', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(jsonData),
  })
  .then(response => response.json())
  .then(data => {

    const cardsContainer = document.getElementById('cards');
    cardsContainer.innerHTML = ''; // Очистить старые карточки

    if (data === null || data.length === 0) {
      cardsContainer.innerHTML = '<p>Проекты не найдены.</p>';
      return;
    }

    data.forEach(item => {
      const card = document.createElement('div');
      card.className = 'col-md-4 mt-4';
      card.innerHTML = `
        <a href="/project/${item.id}" data-bs-toggle="modal" data-bs-target="#exampleModal" onclick="loadProjectData(${item.id})">
          <div class="card p-3">
            <h5>${item.name}</h5>
            <p></p>
            <h3 class="mt-2 primary-color"><i class="bi bi-arrow-right"></i></h3>
          </div>
        </a>
      `;
      cardsContainer.appendChild(card);
    });
  })
  .catch(error => {
    console.error('Ошибка при отправке:', error);
  });
}

  function debounceSend() {
    clearTimeout(debounceTimeout);
    debounceTimeout = setTimeout(() => {
      sendFiltersToAPI();
    }, 500); // Задержка в 500 мс
  }

  document.querySelectorAll('.filter-btn').forEach(button => {
    button.addEventListener('click', () => {
      const id = button.dataset.id;
      const type = button.getAttribute('type'); // "theme" или "specialization"

      button.classList.toggle('active');


      if (button.classList.contains('active')) {
        selectedFilters[type].add(parseInt(id, 10));
        pagination.hidden = true;
      } else {
        selectedFilters[type].delete(parseInt(id, 10));
        if (selectedFilters.specialization.size === 0 && selectedFilters.theme.size === 0) {
          pagination.hidden = false;
        }
      }

      debounceSend(); // Отправить запрос с задержкой
    });
  });

function loadProjectData(projectId) {
    fetch(`/api/project/${projectId}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка при загрузке данных');
            }
            return response.json();
        })
        .then(data => {
            document.getElementById('project-name').textContent = data.name || 'Без названия';

            thematics = data.thematics || [];
            res = "";
            thematics.forEach(theme => {
                res += theme.name + ', ';
            });
            document.getElementById('theme').textContent = res.slice(0, -2) || 'Нет данных';



            document.getElementById('relevance').textContent = data.relevance || 'Нет данных';
            document.getElementById('purpose').textContent = data.purpose || 'Нет данных';
            document.getElementById('result').textContent = data.result || 'Нет данных';

            const directions = data.specializations || [];
            const ul = document.getElementById('specializations');
            ul.innerHTML = ''; 
            directions.forEach(dir => {
                const li = document.createElement('li');
                li.textContent = dir.name;
                ul.appendChild(li);
            });

            const modal = new bootstrap.Modal(document.getElementById('projectModal'));
            modal.show();
        })
        .catch(error => {
            console.error('Ошибка загрузки данных проекта:', error);
        });
}