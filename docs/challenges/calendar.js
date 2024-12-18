const calendarBody = document.getElementById('calendar-body');
const calendarHeader = document.getElementById('calendar-header');
const prevMonthButton = document.getElementById('prev-month');
const nextMonthButton = document.getElementById('next-month');

let currentYear = new Date().getFullYear();
let currentMonth = new Date().getMonth();

function formatDateForSort(day) {
    const [dd, mm, yyyy] = day.split('-');
    return `${yyyy}-${mm.padStart(2, '0')}-${dd.padStart(2, '0')}`;
}

function populateCalendar() {
    calendarBody.innerHTML = '';

    const today = new Date();
    const isCurrentMonth = currentYear === today.getFullYear() && currentMonth === today.getMonth();
    const monthNames = [
        "January", "February", "March", "April", "May", "June", 
        "July", "August", "September", "October", "November", "December"
    ];
    calendarHeader.textContent = `${monthNames[currentMonth]} ${currentYear}`;

    const daysInMonth = new Date(currentYear, currentMonth + 1, 0).getDate();
    const firstDayOfMonth = new Date(currentYear, currentMonth, 1).getDay();

    fetch('data/challenges.json')
        .then(response => {
            if (!response.ok) {
                console.error('Error fetching challenges:', response);
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            data.sort((a, b) => {
                const dateA = new Date(formatDateForSort(a.day));
                const dateB = new Date(formatDateForSort(b.day));
                return dateB - dateA;
            });

            let date = 1;
            for (let i = 0; i < 6; i++) {
                const row = document.createElement('tr');
                for (let j = 0; j < 7; j++) {
                    const cell = document.createElement('td');
                    if (i === 0 && j < firstDayOfMonth) {
                        row.appendChild(cell);
                    } else if (date > daysInMonth) {
                        break;
                    } else {
                        cell.textContent = date;
                        cell.style.textAlign = 'center';

                        const formattedDate = `${String(date).padStart(2, '0')}-${String(currentMonth + 1).padStart(2, '0')}-${currentYear}`;
                        const problem = data.find(p => p.day.startsWith(formattedDate));

                        if (problem) {
                            // cell.style.backgroundColor = "#90ee90";
                            cell.style.cursor = "pointer";

                            cell.addEventListener("click", () => {
                                window.open(problem.url, "_blank");
                            });
                        }

                        if (isCurrentMonth && date === today.getDate()) {
                            cell.style.backgroundColor = "#FFD700";
                            cell.style.borderRadius = "50%";
                            cell.style.color = "#fff";
                        }

                        row.appendChild(cell);
                        date++;
                    }
                }
                calendarBody.appendChild(row);
            }
        })
        .catch(error => {
            console.error('Error loading challenges:', error);
        });
}

prevMonthButton.addEventListener('click', () => {
    currentMonth--;
    if (currentMonth < 0) {
        currentMonth = 11;
        currentYear--;
    }
    populateCalendar();
});

nextMonthButton.addEventListener('click', () => {
    currentMonth++;
    if (currentMonth > 11) {
        currentMonth = 0;
        currentYear++;
    }
    populateCalendar();
});

document.addEventListener('DOMContentLoaded', populateCalendar);
