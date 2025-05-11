document.querySelectorAll('.filter-btn').forEach(button => {

    button.addEventListener('click', () => {
      button.classList.toggle('active');
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