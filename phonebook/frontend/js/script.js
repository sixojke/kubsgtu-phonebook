// В файле js/script.js

const API_BASE_URL = 'http://127.0.0.1:8002';

// Функция сортировки контактов (вынесена отдельно для повторного использования)
function sortContacts(contacts) {
    return [...contacts].sort((a, b) => {
        const hasPhoneA = a.phone ? 1 : 0;
        const hasPhoneB = b.phone ? 1 : 0;
        return hasPhoneB - hasPhoneA; // Сначала те, у кого есть телефон
    });
}

function renderContacts(contacts) {
    // Сортируем контакты перед отображением
    const sortedContacts = sortContacts(contacts);
    
    const list = document.getElementById('contactsList');
    list.innerHTML = sortedContacts.map(contact => `
        <div class="contact-card">
            <div class="contact-icon">
                <i class="fas fa-user-tie"></i>
            </div>
            <div class="contact-info">
                <div class="contact-name">${contact.name}</div>
                <div class="contact-position">${contact.position}</div>
                ${contact.phone ? `<div><i class="fas fa-phone"></i> ${contact.phone}</div>` : ''}
                ${contact.email ? `<div><i class="fas fa-envelope"></i> ${contact.email}</div>` : ''}
                ${contact.room ? `<div><i class="fas fa-door-open"></i> Кабинет: ${contact.room}</div>` : ''}
                <div style="margin-top: 0.5rem; color: var(--primary-blue);">
                    <small>${contact.department}</small>
                </div>
            </div>
        </div>
    `).join('');
}

// Получение отделов
async function loadDepartments() {
    try {
        const response = await fetch(`${API_BASE_URL}/api/departments/list`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
        });
        const data = await response.json();
        return data.departments || [];
    } catch (error) {
        console.error('Ошибка загрузки отделов:', error);
        return [];
    }
}

// Получение контактов с фильтрами
async function loadContacts(page = 1, limit = 100, search = '', department = '') {
    try {
        const response = await fetch(`${API_BASE_URL}/api/contacts/list`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                page,
                limit,
                search,
                department
            })
        });
        return await response.json();
    } catch (error) {
        console.error('Ошибка загрузки контактов:', error);
        return { contacts: [], total: 0 };
    }
}

// Инициализация приложения
document.addEventListener('DOMContentLoaded', async () => {
    // Загружаем отделы для фильтров
    const departments = await loadDepartments();
    renderDepartmentFilters(departments);
    
    // Загружаем контакты
    const { contacts, total } = await loadContacts();
    renderContacts(contacts);
});

// Рендер фильтров отделов
function renderDepartmentFilters(departments) {
    const filterContainer = document.querySelector('.department-filter');
    filterContainer.innerHTML = '';
    
    // Кнопка "Все отделы"
    const allBtn = document.createElement('button');
    allBtn.className = 'filter-btn active';
    allBtn.textContent = 'Все отделы';
    allBtn.addEventListener('click', async () => {
        document.querySelectorAll('.filter-btn').forEach(b => b.classList.remove('active'));
        allBtn.classList.add('active');
        const { contacts } = await loadContacts(1, 100, document.getElementById('searchInput').value, '');
        renderContacts(contacts);
    });
    filterContainer.appendChild(allBtn);
    
    // Кнопки для каждого отдела
    departments.forEach(dept => {
        const btn = document.createElement('button');
        btn.className = 'filter-btn';
        btn.textContent = dept.name;
        btn.addEventListener('click', async () => {
            document.querySelectorAll('.filter-btn').forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            const { contacts } = await loadContacts(
                1, 
                100, 
                document.getElementById('searchInput').value, 
                dept.name
            );
            renderContacts(contacts);
        });
        filterContainer.appendChild(btn);
    });
};

// Поиск с debounce
let searchTimeout;
document.getElementById('searchInput').addEventListener('input', (e) => {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(async () => {
        const activeFilter = document.querySelector('.filter-btn.active').textContent;
        const department = activeFilter === 'Все отделы' ? '' : activeFilter;
        const { contacts } = await loadContacts(1, 100, e.target.value, department);
        renderContacts(contacts);
    }, 300);
});