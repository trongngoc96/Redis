module.exports = {
    apps: [
        {
            name: 'Ewallet-API',
            script: 'app.js',
            node_args: '--max_old_space_size=2048',
            // max_memory_restart : '512M',
            args: '--max_old_space_size=8192',
            autorestart: false,
            instance_var: 'INSTANCE_ID',
            instances: 1,
            increment_var: 'APP_PORT',
            exec_mode: 'fork',
            env: {

                APP_HOST: '0.0.0.0',
                APP_PORT: 8001,
                MINER_RPC_URL: 'http://10.22.7.107:8510',
                MINER_IPC_PATH: '',
                MINER_COINBASE: '0x8d5ee4b23382d7492355bf5c94bc2ee5311b90f6',
                DB_HOST: '10.22.7.107',
                DB_PORT: 3306,
                DB_USER: 'root',
                DB_PASSWORD: 'vnpay123',
                DB_NAME: 'keyman',
                DB_CONNECTION_LIMIT: 20,
                REDIS_HOST: '10.22.7.107'  ,
                REDIS_PORT: 6379,
                REDIS_PASSWORD: '',
                RABBITMQ_URL: 'amqp://vnpay:vnpay123@10.22.7.107',
                API_URL: 'http://127.0.0.1:8001',
            },
            log_type: 'json',
            error_file: '/var/log/ewallet/api_err.log',
            out_file: '/var/log/ewallet/api_out.log',
            log_file: '/var/log/ewallet/api_app.log',
            log_date_format: 'YYYY-MM-DD HH:mm Z',
            merge_logs: true
        }
    ]
};