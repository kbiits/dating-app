const data = [];
const districts = [
    '310101',
    '310102',
    '317101',
    '317102',
    '317103',
    '317104',
    '317105',
    '317106',
    '317107',
    '317108',
    '317201',
    '317202',
    '317203',
    '317204',
    '317205',
    '317206',
    '317301',
    '317302',
    '317303',
    '317304',
    '317305',
    '317306',
    '317307',
    '317308',
    '317401',
    '317402',
    '317403',
    '317404',
    '317405',
    '317406',
    '317407',
    '317408',
    '317409',
];

for (let i = 1; i <= 10; i++) {
    data.push({
        user_name: `User ${i}`,
        user_email: `integration_test_user_${i}@local.mail`,
        user_password: `integration_test_password_${i}`,
        user_bio: `Integration test user ${i} bio`,
        user_gender: i % 2 === 0 ? 'male' : 'female',
        user_dob: '2000-01-01T00:00:00+07:00',
        user_profile_pic_url: 'https://www.example.com/profile_pic.jpg',
        user_district_id: districts[(i - 1) % districts.length],
    });
}

for (let i = 0; i < data.length; i++) {
    data[i].swipes = [];
    // swipe all other users
    const otherUsers = data.slice(0, i).concat(data.slice(i + 1));
    for (let j = 0; j < otherUsers.length; j++) {
        data[i].swipes.push({
            target_user_email: otherUsers[j].user_email,
            target_user_name: otherUsers[j].user_name,
            is_liked: (j + 1) % 2 === 0,
        });
    }
}

console.log(JSON.stringify(data));
