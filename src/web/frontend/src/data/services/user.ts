export class UserServiceClient {
    private static instance: UserServiceClient;

    private client = new proto  

    private constructor() { 

    }

    public static getInstance(): UserServiceClient {
        if (!UserServiceClient.instance) {
            UserServiceClient.instance = new UserServiceClient();
        }

        return UserServiceClient.instance;
    }

    public signUp() {
        // ...
    }
}