import { Request, UnaryResponse } from 'grpc-web';
import { Message } from 'google-protobuf';
import store from "@/store";

class AuthInterceptor {
  intercept(
    request: Request<Message, Message>,
    invoker: (request: Request<Message, Message>) => Promise<UnaryResponse<Message, Message>>
  ): Promise<UnaryResponse<Message, Message>> {
    const token = store.state.auth.accessToken;
    if (token) {
      const metadata = request.getMetadata();
      metadata.Authorization = token;
    }

    return invoker(request);
  }
}

export default AuthInterceptor;
