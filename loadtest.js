import http from 'k6/http';

export let options = {
    scenarios: {
        constant_request_rate: {
            executor: 'constant-arrival-rate',
            rate: 1,
            timeUnit: '1s',
            duration: '10s',
            preAllocatedVUs: 20,
            maxVUs: 30,
        },
    },
};

export default function () {
    let courseNames = ['Math101', 'Eng201'];
    let userId = `user_${__VU}_${__ITER}`;
    let courseName = courseNames[Math.floor(Math.random() * courseNames.length)];

    let payload = JSON.stringify({
        user_id: userId,
        course_name: courseName,
        student_name: `Student ${userId}`,
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post('http://localhost:8080/api/register', payload, params);
}