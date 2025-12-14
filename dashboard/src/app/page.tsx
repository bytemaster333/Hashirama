"use client";

import { useEffect, useState } from "react";
import { Plus, RefreshCw, Server, Trash2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
}
  from "@/components/ui/table";

interface MadaraChain {
  metadata: {
    name: string;
    creationTimestamp: string;
  };
  spec: {
    chainID: string;
    replicas: number;
    port: number;
    image?: string;
  };
  status?: {
    nodesRunning?: number;
  };
}

export default function Dashboard() {
  const [chains, setChains] = useState<MadaraChain[]>([]);
  const [loading, setLoading] = useState(true);
  const [open, setOpen] = useState(false);
  const [formData, setFormData] = useState({
    name: "",
    chainID: "",
    replicas: 1,
  });

  const fetchChains = async () => {
    setLoading(true);
    try {
      const res = await fetch("/api/madarachains");
      if (res.ok) {
        const data = await res.json();
        setChains(data);
      }
    } catch (error) {
      console.error("Failed to fetch chains", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchChains();
  }, []);

  const handleCreate = async () => {
    try {
      const res = await fetch("/api/madarachains", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });
      if (res.ok) {
        setOpen(false);
        setFormData({ name: "", chainID: "", replicas: 1 });
        fetchChains();
      } else {
        alert("Failed to create chain");
      }
    } catch (error) {
      console.error("Failed to create chain", error);
    }
  };

  const handleDelete = async (name: string) => {
    if (!confirm(`Are you sure you want to delete ${name}?`)) return;
    try {
      const res = await fetch(`/api/madarachains?name=${name}`, {
        method: "DELETE",
      });
      if (res.ok) {
        fetchChains();
      } else {
        alert("Failed to delete chain");
      }
    } catch (error) {
      console.error("Failed to delete chain", error);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <div className="mx-auto max-w-6xl space-y-8">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight text-gray-900">
              Hashirama Dashboard
            </h1>
            <p className="text-gray-500">
              Manage your Madara L3 chains on Kubernetes
            </p>
          </div>
          <div className="flex gap-2">
            <Button variant="outline" onClick={fetchChains} disabled={loading}>
              <RefreshCw className={`mr-2 h-4 w-4 ${loading ? "animate-spin" : ""}`} />
              Refresh
            </Button>
            <Dialog open={open} onOpenChange={setOpen}>
              <DialogTrigger asChild>
                <Button>
                  <Plus className="mr-2 h-4 w-4" />
                  Create Chain
                </Button>
              </DialogTrigger>
              <DialogContent>
                <DialogHeader>
                  <DialogTitle>Create New Madara Chain</DialogTitle>
                  <DialogDescription>
                    Deploy a new Starknet L3 sequencer node.
                  </DialogDescription>
                </DialogHeader>
                <div className="grid gap-4 py-4">
                  <div className="grid grid-cols-4 items-center gap-4">
                    <Label htmlFor="name" className="text-right">
                      Name
                    </Label>
                    <Input
                      id="name"
                      value={formData.name}
                      onChange={(e) =>
                        setFormData({ ...formData, name: e.target.value })
                      }
                      className="col-span-3"
                      placeholder="my-chain"
                    />
                  </div>
                  <div className="grid grid-cols-4 items-center gap-4">
                    <Label htmlFor="chainID" className="text-right">
                      Chain ID
                    </Label>
                    <Input
                      id="chainID"
                      value={formData.chainID}
                      onChange={(e) =>
                        setFormData({ ...formData, chainID: e.target.value })
                      }
                      className="col-span-3"
                      placeholder="SN_L3_..."
                    />
                  </div>
                  <div className="grid grid-cols-4 items-center gap-4">
                    <Label htmlFor="replicas" className="text-right">
                      Replicas
                    </Label>
                    <Input
                      id="replicas"
                      type="number"
                      value={formData.replicas}
                      onChange={(e) =>
                        setFormData({
                          ...formData,
                          replicas: parseInt(e.target.value),
                        })
                      }
                      className="col-span-3"
                      min={1}
                    />
                  </div>
                </div>
                <DialogFooter>
                  <Button onClick={handleCreate}>Create</Button>
                </DialogFooter>
              </DialogContent>
            </Dialog>
          </div>
        </div>

        <Card>
          <CardHeader>
            <CardTitle>Active Chains</CardTitle>
            <CardDescription>
              List of deployed MadaraChain resources.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Name</TableHead>
                  <TableHead>Chain ID</TableHead>
                  <TableHead>Replicas</TableHead>
                  <TableHead>Port</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead>Age</TableHead>
                  <TableHead className="text-right">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {chains.length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={7} className="text-center h-24 text-gray-500">
                      No chains found. Create one to get started.
                    </TableCell>
                  </TableRow>
                ) : (
                  chains.map((chain) => (
                    <TableRow key={chain.metadata.name}>
                      <TableCell className="font-medium flex items-center gap-2">
                        <Server className="h-4 w-4 text-blue-500" />
                        {chain.metadata.name}
                      </TableCell>
                      <TableCell>{chain.spec.chainID}</TableCell>
                      <TableCell>{chain.spec.replicas}</TableCell>
                      <TableCell>{chain.spec.port}</TableCell>
                      <TableCell>
                        <span className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${(chain.status?.nodesRunning || 0) === chain.spec.replicas
                          ? "bg-green-100 text-green-800"
                          : "bg-yellow-100 text-yellow-800"
                          }`}>
                          {chain.status?.nodesRunning || 0} / {chain.spec.replicas} Running
                        </span>
                      </TableCell>
                      <TableCell>
                        {new Date(chain.metadata.creationTimestamp).toLocaleDateString()}
                      </TableCell>
                      <TableCell className="text-right">
                        <Button
                          variant="ghost"
                          size="icon"
                          className="text-red-500 hover:text-red-700 hover:bg-red-50"
                          onClick={() => handleDelete(chain.metadata.name)}
                        >
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </TableCell>
                    </TableRow>
                  ))
                )}
              </TableBody>
            </Table>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
