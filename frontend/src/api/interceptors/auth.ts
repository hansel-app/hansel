import { ClientReadableStream, Request, UnaryResponse } from 'grpc-web';
import { Message } from 'google-protobuf';

class AuthInterceptor {
  private token: string;

  constructor(token: string) {
    this.token = token;
  }

  intercept(
    request: Request<Message, Message>,
    invoker: (request: Request<Message, Message>) => UnaryResponse<Message, Message> | ClientReadableStream<Message>
  ): UnaryResponse<Message, Message> | ClientReadableStream<Message> {
    const metadata = request.getMetadata()
    metadata.Authorization = 'Bearer ' + this.token
    return invoker(request)
  }
}

export default AuthInterceptor;
