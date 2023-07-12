package graphql

import (
	"cinematic.back/api/users"
	"cinematic.back/api/users/pb"
	"cinematic.back/services/gateway/internal/api/graphql/generated"
	"cinematic.back/services/gateway/internal/api/graphql/resolvers"
	"cinematic.back/services/gateway/internal/services/user"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"net/http"
	"strings"
	"time"
)

const endpointUrl = "/graphql"

type Server struct {
	uClient  users.Client
	uService user.Service
	handler  *handler.Server
}

func NewServer(
	uClient users.Client,
	uService user.Service,
) *Server {
	s := &Server{uClient: uClient, uService: uService}

	cfg := generated.Config{
		Resolvers: resolvers.NewResolver(uClient, uService),
	}

	s.setDirectives(&cfg)

	h := handler.New(generated.NewExecutableSchema(cfg))

	h.AddTransport(transport.Websocket{KeepAlivePingInterval: 10 * time.Second})
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.Use(extension.Introspection{})

	s.handler = h

	return s
}

func (s *Server) authMiddleware(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tokens := strings.Split(request.Header.Get("Authorization"), "Bearer ")
		if len(tokens) < 2 {
			next.ServeHTTP(writer, request)
			return
		}

		token := tokens[1]

		auth, err := s.uClient.AuthByAccessToken(request.Context(), &pb.AuthByAccessTokenRequest{Token: token})
		if err != nil {
			next.ServeHTTP(writer, request)
			return
		}

		ctx := s.uService.ToOutgoingCtx(request.Context(), auth.User)

		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	}
}

func (s *Server) setDirectives(cfg *generated.Config) {
	cfg.Directives.IsLogged = func(
		ctx context.Context,
		obj interface{},
		next graphql.Resolver,
	) (res interface{}, err error) {
		if _, ok := s.uService.FromIncomingCtx(ctx); ok {
			next(ctx)
		}

		return nil, err
	}
}

func (s *Server) Start() error {
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		Debug:            true,
	}).Handler)

	h := s.authMiddleware(s.handler)

	router.Handle("/graphiql", playground.Handler("GraphQL playground", endpointUrl))
	router.Handle(endpointUrl, h)

	return http.ListenAndServe(":3000", router)
}
