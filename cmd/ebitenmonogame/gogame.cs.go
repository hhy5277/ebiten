// Code generated by file2byteslice. DO NOT EDIT.
// (gofmt is fine after generating)

package main

var gogame_cs = []byte("\ufeffusing Microsoft.Xna.Framework;\nusing Microsoft.Xna.Framework.Graphics;\nusing Microsoft.Xna.Framework.Input;\nusing System;\nusing System.IO;\n\nusing {{.Namespace}}.AutoGen;\n\nnamespace {{.Namespace}}\n{\n    // Operation must sync with driver.Operation.\n    enum Operation\n    {\n        Zero,\n        One,\n        SrcAlpha,\n        DstAlpha,\n        OneMinusSrcAlpha,\n        OneMinusDstAlpha,\n    }\n\n    public class GoGame : Game\n    {\n        private GraphicsDeviceManager graphics;\n        private IInvokable onUpdate;\n        private IInvokable onDraw;\n        private VertexBuffer vertexBuffer;\n        private IndexBuffer indexBuffer;\n        private Effect effect;\n\n        public GoGame(IInvokable onUpdate, IInvokable onDraw)\n        {\n            this.onUpdate = onUpdate;\n            this.onDraw = onDraw;\n            this.graphics = new GraphicsDeviceManager(this);\n            this.graphics.PreferredBackBufferWidth = 640;\n            this.graphics.PreferredBackBufferHeight = 480;\n\n            this.Content.RootDirectory = \"Content\";\n            this.IsMouseVisible = true;\n        }\n\n        protected override void LoadContent()\n        {\n            VertexElement[] elements = new VertexElement[]\n            {\n                new VertexElement(sizeof(float)*0, VertexElementFormat.Vector2, VertexElementUsage.Position, 0),\n                new VertexElement(sizeof(float)*2, VertexElementFormat.Vector2, VertexElementUsage.TextureCoordinate, 0),\n                new VertexElement(sizeof(float)*4, VertexElementFormat.Vector4, VertexElementUsage.TextureCoordinate, 1),\n                new VertexElement(sizeof(float)*8, VertexElementFormat.Vector4, VertexElementUsage.Color, 0),\n            };\n            this.vertexBuffer = new DynamicVertexBuffer(\n                this.GraphicsDevice, new VertexDeclaration(elements), 65536, BufferUsage.None);\n            this.GraphicsDevice.SetVertexBuffer(this.vertexBuffer);\n\n            this.indexBuffer = new DynamicIndexBuffer(\n                this.GraphicsDevice, IndexElementSize.SixteenBits, 65536, BufferUsage.None);\n            this.GraphicsDevice.Indices = this.indexBuffer;\n            \n            // TODO: Add more shaders for e.g., linear filter.\n            this.effect = Content.Load<Effect>(\"Shader\");\n\n            this.GraphicsDevice.RasterizerState = new RasterizerState()\n            {\n                CullMode = CullMode.None,\n            };\n\n            base.LoadContent();\n        }\n\n        internal void SetDestination(RenderTarget2D renderTarget2D, int viewportWidth, int viewportHeight)\n        {\n            this.GraphicsDevice.SetRenderTarget(renderTarget2D);\n            this.GraphicsDevice.Viewport = new Viewport(0, 0, viewportWidth, viewportHeight);\n            this.effect.Parameters[\"ViewportSize\"].SetValue(new Vector2(viewportWidth, viewportHeight));\n        }\n\n        internal void SetSource(Texture2D texture2D)\n        {\n            this.effect.Parameters[\"Texture\"].SetValue(texture2D);\n        }\n\n        internal void SetVertices(byte[] vertices, byte[] indices)\n        {\n            this.vertexBuffer.SetData(vertices, 0, vertices.Length);\n            this.indexBuffer.SetData(indices, 0, indices.Length);\n        }\n\n        protected override void Update(GameTime gameTime)\n        {\n            this.onUpdate.Invoke(null);\n            base.Update(gameTime);\n        }\n\n        protected override void Draw(GameTime gameTime)\n        {\n            this.onDraw.Invoke(null);\n            base.Draw(gameTime);\n        }\n\n        internal void DrawTriangles(int indexLen, int indexOffset, Blend blendSrc, Blend blendDst)\n        {\n            this.GraphicsDevice.BlendState = new BlendState()\n            {\n                AlphaSourceBlend = blendSrc,\n                ColorSourceBlend = blendSrc,\n                AlphaDestinationBlend = blendDst,\n                ColorDestinationBlend = blendDst,\n            };\n            foreach (EffectPass pass in this.effect.CurrentTechnique.Passes)\n            {\n                pass.Apply();\n                this.GraphicsDevice.DrawIndexedPrimitives(\n                    PrimitiveType.TriangleList, 0, indexOffset, indexLen / 3);\n            }\n        }\n    }\n\n    // This methods are called from Go world. They must not have overloads.\n    class GameGoBinding\n    {\n        private static Blend ToBlend(Operation operation)\n        {\n            switch (operation)\n            {\n                case Operation.Zero:\n                    return Blend.Zero;\n                case Operation.One:\n                    return Blend.One;\n                case Operation.SrcAlpha:\n                    return Blend.SourceAlpha;\n                case Operation.DstAlpha:\n                    return Blend.DestinationAlpha;\n                case Operation.OneMinusSrcAlpha:\n                    return Blend.InverseSourceAlpha;\n                case Operation.OneMinusDstAlpha:\n                    return Blend.InverseDestinationAlpha;\n            }\n            throw new ArgumentOutOfRangeException(\"operation\", operation, \"\");\n        }\n\n        private GoGame game;\n\n        private GameGoBinding(IInvokable onUpdate, IInvokable onDraw)\n        {\n            this.game = new GoGame(onUpdate, onDraw);\n        }\n\n        private void Run()\n        {\n            try\n            {\n                this.game.Run();\n            }\n            finally\n            {\n                this.game.Dispose();\n            }\n        }\n\n        private RenderTarget2D NewRenderTarget2D(double width, double height)\n        {\n            return new RenderTarget2D(this.game.GraphicsDevice, (int)width, (int)height);\n        }\n\n        private void Dispose(Texture2D texture2D)\n        {\n            texture2D.Dispose();\n        }\n\n        private void ReplacePixels(RenderTarget2D renderTarget2D, byte[] pixels, double x, double y, double width, double height)\n        {\n            var rect = new Rectangle((int)x, (int)y, (int)width, (int)height);\n            renderTarget2D.SetData(0, rect, pixels, 0, pixels.Length);\n        }\n\n        private byte[] Pixels(Texture2D texture2D, double width, double height)\n        {\n            Rectangle rect = new Rectangle(0, 0, (int)width, (int)height);\n            byte[] data = new byte[4*(int)width*(int)height];\n            texture2D.GetData(0, rect, data, 0, data.Length);\n            return data;\n        }\n\n        private void SetDestination(RenderTarget2D renderTarget2D, double viewportWidth, double viewportHeight)\n        {\n            this.game.SetDestination(renderTarget2D, (int)viewportWidth, (int)viewportHeight);\n        }\n\n        private void SetSource(Texture2D texture2D)\n        {\n            this.game.SetSource(texture2D);\n        }\n\n        private void SetVertices(byte[] vertices, byte[] indices)\n        {\n            this.game.SetVertices(vertices, indices);\n        }\n\n        private void Draw(double indexLen, double indexOffset, double operationSrc, double operationDst)\n        {\n            this.game.DrawTriangles((int)indexLen, (int)indexOffset, ToBlend((Operation)operationSrc), ToBlend((Operation)operationDst));\n        }\n\n        private bool IsKeyPressed(string driverKey)\n        {\n            Keys key;\n            switch (driverKey)\n            {\n                case \"KeyDown\":\n                    key = Keys.Down;\n                    break;\n                case \"KeyLeft\":\n                    key = Keys.Left;\n                    break;\n                case \"KeySpace\":\n                    key = Keys.Space;\n                    break;\n                case \"KeyRight\":\n                    key = Keys.Right;\n                    break;\n                case \"KeyUp\":\n                    key = Keys.Up;\n                    break;\n                default:\n                    return false;\n            }\n            return Keyboard.GetState().IsKeyDown(key);\n        }\n    }\n}\n")
