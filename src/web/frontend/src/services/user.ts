export class UserServiceClient {
  private static instance: UserServiceClient;

  private constructor() {}

  public static getInstance(): UserServiceClient {
    if (!UserServiceClient.instance) {
      UserServiceClient.instance = new UserServiceClient();
    }

    return UserServiceClient.instance;
  }

  public async signUp() {
    // ...
    const response = await fetch("/api");
    const data = response.json()
  }
}
