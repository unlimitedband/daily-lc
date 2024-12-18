function formatDate(customDateString) {
    try {
        // Chuyển đổi định dạng "18-12-2024 01:55 UTC" thành đối tượng Date
        const [day, month, year, time, zone] = customDateString.split(/[- :]/);
        const isoDate = `${year}-${month}-${day}T${time}:00Z`; // Định dạng ISO 8601
        const date = new Date(isoDate);

        // Format ngày cho đẹp mắt
        const options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };
        return date.toLocaleDateString('en-US', options);
    } catch (error) {
        console.error('Error parsing date:', error);
        return customDateString; // Trả về ngày gốc nếu lỗi
    }
}

function formatDifficulty(difficulty) {
    return difficulty.charAt(0).toUpperCase() + difficulty.slice(1).toLowerCase(); // Viết hoa chữ đầu
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
                const dateA = new Date(formatDateForSort(a.day));
                const dateB = new Date(formatDateForSort(b.day));
                return dateB - dateA; // Ngày mới nhất lên trước
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

// Chuyển đổi định dạng "day" cho mục đích sắp xếp
function formatDateForSort(customDateString) {
    const [day, month, year, time] = customDateString.split(/[- :]/);
    return `${year}-${month}-${day}T${time}:00Z`; // ISO 8601
}

document.addEventListener('DOMContentLoaded', loadChallenges);
