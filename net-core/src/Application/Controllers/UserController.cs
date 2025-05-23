namespace Boilerplate.Controllers;

using Application.DTOs.User;
using Application.Middlewares;
using Application.Service;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("/api/v1/users")]
public class UserController : ControllerBase
{

  private readonly ILogger<UserController> _logger;
  private readonly IUserService _service;

  public UserController(ILogger<UserController> logger, IUserService service)
  {
    _logger = logger;
    _service = service;
  }

  [Protected]
  [AdminOnly]
  [HttpGet("admin")]
  public async Task<IActionResult> Get([FromQuery] UserQuery query)
  {
    _logger.LogInformation("Get all users request received");

    return await _service.HandleGetAllAsync(query);
  }

  [Protected]
  [AdminOnly]
  [HttpPost("admin")]
  public async Task<IActionResult> Create([FromBody] UserCreateDto req)
  {
    _logger.LogInformation("Create user request received");

    return await _service.HandleCreateAsync(req);
  }

  [Protected]
  [AdminOnly]
  [HttpPut("admin/{id}")]
  public async Task<IActionResult> Update(Guid id, [FromBody] UserUpdateDto req)
  {
    _logger.LogInformation("Update user request received");

    return await _service.HandleUpdateAsync(id, req);
  }

  [Protected]
  [HttpGet("profile")]
  public async Task<IActionResult> GetProfile()
  {
    _logger.LogInformation("Get user profile request received");

    return await _service.HandleGetProfileAsync();
  }

  [Protected]
  [HttpPut("profile")]
  public async Task<IActionResult> UpdateProfile([FromBody] UpdateProfileReq req)
  {
    _logger.LogInformation("Update user profile request received");

    return await _service.HandleUpdateProfileAsync(req);
  }
}
