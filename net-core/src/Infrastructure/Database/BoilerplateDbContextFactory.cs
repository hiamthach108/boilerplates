namespace Database;

using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Design;

public class BoilerplateDbContextFactory : IDesignTimeDbContextFactory<BoilerplateDbContext>
{
  public BoilerplateDbContext CreateDbContext(string[] args)
  {
    var configuration = new ConfigurationBuilder()
      .AddJsonFile("appsettings.json")
      .Build();

    Console.WriteLine($"Using ConnectionString: {configuration.GetConnectionString("DatabaseConnection")}");

    var optionsBuilder = new DbContextOptionsBuilder<BoilerplateDbContext>();
    optionsBuilder.UseNpgsql(
      configuration.GetConnectionString("DatabaseConnection") ?? "",
      x => x.UseNetTopologySuite()
    );

    return new BoilerplateDbContext(optionsBuilder.Options);
  }
}
