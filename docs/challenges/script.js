function formatDate(customDateString) {
    try {
        const [day, month, year, time] = customDateString.split(/[- :]/);
        const isoDate = `${year}-${month}-${day}T${time}:00Z`;
        const date = new Date(isoDate);

        const options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };
        return date.toLocaleDateString('en-US', options);
    } catch (error) {
        console.error('Error parsing date:', error);
        return customDateString;
    }
}

function formatDifficulty(difficulty) {
    return difficulty.charAt(0).toUpperCase() + difficulty.slice(1).toLowerCase();
}

function loadChallenges() {
    fetch('data/challenges.json')
        .then(response => {
            if (!response.ok) {
                console.error('Error fetching challenges:', response);
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            // Sắp xếp các challenges theo ngày mới nhất
            data.sort((a, b) => {
                const dateA = new Date(a.day);
                const dateB = new Date(b.day);
                return dateB - dateA; // Ngày mới nhất sẽ được ưu tiên
            });

            const problemList = document.getElementById('problemList');
            problemList.innerHTML = '';

            data.forEach(challenge => {
                const formattedDate = formatDate(challenge.day);
                const formattedDifficulty = formatDifficulty(challenge.difficulty);

                const listItem = document.createElement('tr');
                listItem.innerHTML = `
                    <td>${formattedDate}</td>
                    <td>${challenge.title}</td>
                    <td>${formattedDifficulty}</td>
                    <td>
                        <button class="arrow-btn" onclick="window.open('${challenge.url}', '_blank')">➡️</button>
                    </td>
                `;
                problemList.appendChild(listItem);
            });
        })
        .catch(error => {
            console.error('Error loading challenges:', error);
        });
}

document.addEventListener('DOMContentLoaded', loadChallenges);
