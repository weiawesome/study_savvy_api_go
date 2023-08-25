import http from 'k6/http';
import { sleep,check } from 'k6';

export let options = {
    stages: [
        { duration: '1m', target: 10 },
        { duration: '1m', target: 1500 },
        { duration: '1m', target: 10 },
    ],
};
export default function () {
    const url = "https://study-savvy.com/api/User/login/app";
    const payload = {
        mail:"open891013@gmail.com",
        password: "Wei891013",
    };

    const headers = { 'Content-Type': 'application/json' };

    const response = http.post(url, JSON.stringify(payload), { headers: headers });

    check(response, {
        'Status is 200': (res) => res.status === 200,
    });
    check(response, {
        'Response body has token property': (res) => JSON.parse(res.body).hasOwnProperty('token'),
    });

    sleep(2);
}