import { Request, UnaryResponse } from 'grpc-web';
import { Message } from 'google-protobuf';
import { LoginResponse } from "@/protobuf/auth_pb";

const AUTHORIZATION_HEADER_PREFIX = 'Bearer ';

class AuthInterceptor {
  private token: string;

  constructor(token?: string) {
    this.token = token ?? "";
  }

  intercept(
    request: Request<Message, Message>,
    invoker: (request: Request<Message, Message>) => Promise<UnaryResponse<Message, Message>>
  ): Promise<UnaryResponse<Message, Message>> {
    const metadata = request.getMetadata();
    metadata.Authorization = AUTHORIZATION_HEADER_PREFIX + this.token;

    return invoker(request).then((response: UnaryResponse<Message, Message>) => {
      const responseMessage = response.getResponseMessage();
      if (responseMessage instanceof LoginResponse) {
        this.token = responseMessage.getAccessToken();
      }
      return response;
    });
  }
}

export default AuthInterceptor;
